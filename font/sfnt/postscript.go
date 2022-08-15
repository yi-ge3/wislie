// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sfnt

// Compact Font Format (CFF) fonts are written in PostScript, a stack-based
// programming language.
//
// A fundamental concept is a DICT, or a key-value map, expressed in reverse
// Polish notation. For example, this sequence of operations:
//	- push the number 379
//	- version operator
//	- push the number 392
//	- Notice operator
//	- etc
//	- push the number 100
//	- push the number 0
//	- push the number 500
//	- push the number 800
//	- FontBBox operator
//	- etc
// defines a DICT that maps "version" to the String ID (SID) 379, the copyright
// "Notice" to the SID 392, the font bounding box "FontBBox" to the four
// numbers [100, 0, 500, 800], etc.
//
// The first 391 String IDs (starting at 0) are predefined as per the CFF spec
// Appendix A, in 5176.CFF.pdf referenced below. For example, 379 means
// "001.000". String ID 392 is not predefined, and is mapped by a separate
// structure, the "String INDEX", inside the CFF data. (String ID 391 is also
// not predefined. Specifically for ../testdata/CFFTest.otf, 391 means
// "uni4E2D", as this font contains a glyph for U+4E2D).
//
// The actual glyph vectors are similarly encoded (in PostScript), in a format
// called Type 2 Charstrings. The wire encoding is similar to but not exactly
// the same as CFF's. For example, the byte 0x05 means FontBBox for CFF DICTs,
// but means rlineto (relative line-to) for Type 2 Charstrings. See
// 5176.CFF.pdf Appendix H and 5177.Type2.pdf Appendix A in the PDF files
// referenced below.
//
// CFF is a stand-alone format, but CFF as used in SFNT fonts have further
// restrictions. For example, a stand-alone CFF can contain multiple fonts, but
// https://www.microsoft.com/typography/OTSPEC/cff.htm says that "The Name
// INDEX in the CFF must contain only one entry; that is, there must be only
// one font in the CFF FontSet".
//
// The relevant specifications are:
// 	- http://wwwimages.adobe.com/content/dam/Adobe/en/devnet/font/pdfs/5176.CFF.pdf
// 	- http://wwwimages.adobe.com/content/dam/Adobe/en/devnet/font/pdfs/5177.Type2.pdf

import (
	"fmt"
)

const (
	// psStackSize is the stack size for a PostScript interpreter. 5176.CFF.pdf
	// section 4 "DICT Data" says that "An operator may be preceded by up to a
	// maximum of 48 operands". Similarly, 5177.Type2.pdf Appendix B "Type 2
	// Charstring Implementation Limits" says that "Argument stack 48".
	psStackSize = 48
)

func bigEndian(b []byte) uint32 {
	switch len(b) {
	case 1:
		return uint32(b[0])
	case 2:
		return uint32(b[0])<<8 | uint32(b[1])
	case 3:
		return uint32(b[0])<<16 | uint32(b[1])<<8 | uint32(b[2])
	case 4:
		return uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
	}
	panic("unreachable")
}

// cffParser parses the CFF table from an SFNT font.
type cffParser struct {
	src    *source
	base   int
	offset int
	end    int
	buf    []byte
	err    error
	locBuf [2]uint32

	instructions []byte

	stack struct {
		a   [psStackSize]int32
		top int32
	}

	saved struct {
		charStrings int32
	}
}

func (p *cffParser) parse() (locations []uint32, err error) {
	// Parse header.
	{
		if !p.read(4) {
			return nil, p.err
		}
		if p.buf[0] != 1 || p.buf[1] != 0 || p.buf[2] != 4 {
			return nil, errUnsupportedCFFVersion
		}
	}

	// Parse Name INDEX.
	{
		count, offSize, ok := p.parseIndexHeader()
		if !ok {
			return nil, p.err
		}
		// https://www.microsoft.com/typography/OTSPEC/cff.htm says that "The
		// Name INDEX in the CFF must contain only one entry".
		if count != 1 {
			return nil, errInvalidCFFTable
		}
		if !p.parseIndexLocations(p.locBuf[:2], count, offSize) {
			return nil, p.err
		}
		p.offset = int(p.locBuf[1])
	}

	// Parse Top DICT INDEX.
	{
		count, offSize, ok := p.parseIndexHeader()
		if !ok {
			return nil, p.err
		}
		// 5176.CFF.pdf section 8 "Top DICT INDEX" says that the count here
		// should match the count of the Name INDEX, which is 1.
		if count != 1 {
			return nil, errInvalidCFFTable
		}
		if !p.parseIndexLocations(p.locBuf[:2], count, offSize) {
			return nil, p.err
		}
		if !p.read(int(p.locBuf[1] - p.locBuf[0])) {
			return nil, p.err
		}

		for p.instructions = p.buf; len(p.instructions) > 0; {
			p.step()
			if p.err != nil {
				return nil, p.err
			}
		}
	}

	// Parse the CharStrings INDEX, whose location was found in the Top DICT.
	if p.saved.charStrings <= 0 || int32(p.end-p.base) < p.saved.charStrings {
		return nil, errInvalidCFFTable
	}
	p.offset = p.base + int(p.saved.charStrings)
	count, offSize, ok := p.parseIndexHeader()
	if !ok {
		return nil, p.err
	}
	if count == 0 {
		return nil, errInvalidCFFTable
	}
	locations = make([]uint32, count+1)
	if !p.parseIndexLocations(locations, count, offSize) {
		return nil, p.err
	}
	return locations, nil
}

// read sets p.buf to view the n bytes from p.offset to p.offset+n. It also
// advances p.offset by n.
//
// As per the source.view method, the caller should not modify the contents of
// p.buf after read returns, other than by calling read again.
//
// The caller should also avoid modifying the pointer / length / capacity of
// the p.buf slice, not just avoid modifying the slice's contents, in order to
// maximize the opportunity to re-use p.buf's allocated memory when viewing the
// underlying source data for subsequent read calls.
func (p *cffParser) read(n int) (ok bool) {
	if p.end-p.offset < n {
		p.err = errInvalidCFFTable
		return false
	}
	p.buf, p.err = p.src.view(p.buf, p.offset, n)
	p.offset += n
	return p.err == nil
}

func (p *cffParser) parseIndexHeader() (count, offSize int32, ok bool) {
	if !p.read(2) {
		return 0, 0, false
	}
	count = int32(u16(p.buf[:2]))
	// 5176.CFF.pdf section 5 "INDEX Data" says that "An empty INDEX is
	// represented by a count field with a 0 value and no additional fields.
	// Thus, the total size of an empty INDEX is 2 bytes".
	if count == 0 {
		return count, 0, true
	}
	if !p.read(1) {
		return 0, 0, false
	}
	offSize = int32(p.buf[0])
	if offSize < 1 || 4 < offSize {
		p.err = errInvalidCFFTable
		return 0, 0, false
	}
	return count, offSize, true
}

func (p *cffParser) parseIndexLocations(dst []uint32, count, offSize int32) (ok bool) {
	if count == 0 {
		return true
	}
	if len(dst) != int(count+1) {
		panic("unreachable")
	}
	if !p.read(len(dst) * int(offSize)) {
		return false
	}

	buf, prev := p.buf, uint32(0)
	for i := range dst {
		loc := bigEndian(buf[:offSize])
		buf = buf[offSize:]

		// Locations are off by 1 byte. 5176.CFF.pdf section 5 "INDEX Data"
		// says that "Offsets in the offset array are relative to the byte that
		// precedes the object data... This ensures that every object has a
		// corresponding offset which is always nonzero".
		if loc == 0 {
			p.err = errInvalidCFFTable
			return false
		}
		loc--

		// In the same paragraph, "Therefore the first element of the offset
		// array is always 1" before correcting for the off-by-1.
		if i == 0 {
			if loc != 0 {
				p.err = errInvalidCFFTable
				break
			}
		} else if loc <= prev { // Check that locations are increasing.
			p.err = errInvalidCFFTable
			break
		}

		// Check that locations are in bounds.
		if uint32(p.end-p.offset) < loc {
			p.err = errInvalidCFFTable
			break
		}

		dst[i] = uint32(p.offset) + loc
		prev = loc
	}
	return p.err == nil
}

// step executes a single operation, whether pushing a numeric operand onto the
// stack or executing an operator.
func (p *cffParser) step() {
	if number, res := p.parseNumber(); res != prNone {
		if res == prBad || p.stack.top == psStackSize {
			p.err = errInvalidCFFTable
			return
		}
		p.stack.a[p.stack.top] = number
		p.stack.top++
		return
	}

	b0 := p.instructions[0]
	p.instructions = p.instructions[1:]
	if int(b0) < len(cff1ByteOperators) {
		if op := cff1ByteOperators[b0]; op.name != "" {
			if p.stack.top < op.numPop {
				p.err = errInvalidCFFTable
				return
			}
			if op.run != nil {
				op.run(p)
			}
			if op.numPop < 0 {
				p.stack.top = 0
			} else {
				p.stack.top -= op.numPop
			}
			return
		}
	}

	p.err = fmt.Errorf("sfnt: unrecognized CFF 1-byte operator %d", b0)
}

type parseResult int32

const (
	prBad  parseResult = -1
	prNone parseResult = +0
	prGood parseResult = +1
)

// See 5176.CFF.pdf section 4 "DICT Data".
func (p *cffParser) parseNumber() (number int32, res parseResult) {
	if len(p.instructions) == 0 {
		return 0, prNone
	}

	switch b0 := p.instructions[0]; {
	case b0 == 28:
		if len(p.instructions) < 3 {
			return 0, prBad
		}
		number = int32(int16(u16(p.instructions[1:])))
		p.instructions = p.instructions[3:]
		return number, prGood

	case b0 == 29:
		if len(p.instructions) < 5 {
			return 0, prBad
		}
		number = int32(u32(p.instructions[1:]))
		p.instructions = p.instructions[5:]
		return number, prGood

	case b0 < 32:
		// No-op.

	case b0 < 247:
		p.instructions = p.instructions[1:]
		return int32(b0) - 139, prGood

	case b0 < 251:
		if len(p.instructions) < 2 {
			return 0, prBad
		}
		b1 := p.instructions[1]
		p.instructions = p.instructions[2:]
		return +int32(b0-247)*256 + int32(b1) + 108, prGood

	case b0 < 255:
		if len(p.instructions) < 2 {
			return 0, prBad
		}
		b1 := p.instructions[1]
		p.instructions = p.instructions[2:]
		return -int32(b0-251)*256 - int32(b1) - 108, prGood
	}

	return 0, prNone
}

type cffOperator struct {
	// numPop is the number of stack values to pop. -1 means "array" and -2
	// means "delta" as per 5176.CFF.pdf Table 6 "Operand Types".
	numPop int32
	// name is the operator name. An empty name (i.e. the zero value for the
	// struct overall) means an unrecognized 1-byte operator.
	name string
	// run is the function that implements the operator. Nil means that we
	// ignore the operator, other than popping its arguments off the stack.
	run func(*cffParser)
}

// cff1ByteOperators encodes the subset of 5176.CFF.pdf Table 9 "Top DICT
// Operator Entries" used by this implementation.
var cff1ByteOperators = [...]cffOperator{
	0:  {+1, "version", nil},
	1:  {+1, "Notice", nil},
	2:  {+1, "FullName", nil},
	3:  {+1, "FamilyName", nil},
	4:  {+1, "Weight", nil},
	5:  {-1, "FontBBox", nil},
	13: {+1, "UniqueID", nil},
	14: {-1, "XUID", nil},
	15: {+1, "charset", nil},
	16: {+1, "Encoding", nil},
	17: {+1, "CharStrings", func(p *cffParser) {
		p.saved.charStrings = p.stack.a[p.stack.top-1]
	}},
	18: {+2, "Private", nil},
}

// TODO: 2-byte operators.
