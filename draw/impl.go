// generated by "go run gen.go". DO NOT EDIT.

package draw

import (
	"image"
	"image/color"
)

func (z *nnScaler) Scale(dst Image, dp image.Point, src image.Image, sp image.Point) {
	if z.dw <= 0 || z.dh <= 0 || z.sw <= 0 || z.sh <= 0 {
		return
	}
	// adr is the affected destination pixels, relative to dp.
	adr := dst.Bounds().Sub(dp).Intersect(image.Rectangle{Max: image.Point{int(z.dw), int(z.dh)}})
	if adr.Empty() {
		return
	}
	// sr is the source pixels. If it extends beyond the src bounds,
	// we cannot use the type-specific fast paths, as they access
	// the Pix fields directly without bounds checking.
	if sr := (image.Rectangle{sp, sp.Add(image.Point{int(z.sw), int(z.sh)})}); !sr.In(src.Bounds()) {
		z.scale_Image_Image(dst, dp, adr, src, sp)
	} else {
		switch dst := dst.(type) {
		case *image.RGBA:
			switch src := src.(type) {
			case *image.Gray:
				z.scale_RGBA_Gray(dst, dp, adr, src, sp)
			case *image.NRGBA:
				z.scale_RGBA_NRGBA(dst, dp, adr, src, sp)
			case *image.RGBA:
				z.scale_RGBA_RGBA(dst, dp, adr, src, sp)
			case *image.Uniform:
				z.scale_RGBA_Uniform(dst, dp, adr, src, sp)
			case *image.YCbCr:
				z.scale_RGBA_YCbCr(dst, dp, adr, src, sp)
			default:
				z.scale_RGBA_Image(dst, dp, adr, src, sp)
			}
		default:
			switch src := src.(type) {
			default:
				z.scale_Image_Image(dst, dp, adr, src, sp)
			}
		}
	}
}

func (z *nnScaler) scale_RGBA_Gray(dst *image.RGBA, dp image.Point, adr image.Rectangle, src *image.Gray, sp image.Point) {
	for dy := int32(adr.Min.Y); dy < int32(adr.Max.Y); dy++ {
		sy := (2*uint64(dy) + 1) * uint64(z.sh) / (2 * uint64(z.dh))
		d := dst.PixOffset(dp.X+adr.Min.X, dp.Y+int(dy))
		for dx := int32(adr.Min.X); dx < int32(adr.Max.X); dx++ {
			sx := (2*uint64(dx) + 1) * uint64(z.sw) / (2 * uint64(z.dw))
			pr, pg, pb, pa := src.At(sp.X+int(sx), sp.Y+int(sy)).RGBA()
			dst.Pix[d+0] = uint8(uint32(pr) >> 8)
			dst.Pix[d+1] = uint8(uint32(pg) >> 8)
			dst.Pix[d+2] = uint8(uint32(pb) >> 8)
			dst.Pix[d+3] = uint8(uint32(pa) >> 8)
			d += 4
		}
	}
}

func (z *nnScaler) scale_RGBA_NRGBA(dst *image.RGBA, dp image.Point, adr image.Rectangle, src *image.NRGBA, sp image.Point) {
	for dy := int32(adr.Min.Y); dy < int32(adr.Max.Y); dy++ {
		sy := (2*uint64(dy) + 1) * uint64(z.sh) / (2 * uint64(z.dh))
		d := dst.PixOffset(dp.X+adr.Min.X, dp.Y+int(dy))
		for dx := int32(adr.Min.X); dx < int32(adr.Max.X); dx++ {
			sx := (2*uint64(dx) + 1) * uint64(z.sw) / (2 * uint64(z.dw))
			pr, pg, pb, pa := src.At(sp.X+int(sx), sp.Y+int(sy)).RGBA()
			dst.Pix[d+0] = uint8(uint32(pr) >> 8)
			dst.Pix[d+1] = uint8(uint32(pg) >> 8)
			dst.Pix[d+2] = uint8(uint32(pb) >> 8)
			dst.Pix[d+3] = uint8(uint32(pa) >> 8)
			d += 4
		}
	}
}

func (z *nnScaler) scale_RGBA_RGBA(dst *image.RGBA, dp image.Point, adr image.Rectangle, src *image.RGBA, sp image.Point) {
	for dy := int32(adr.Min.Y); dy < int32(adr.Max.Y); dy++ {
		sy := (2*uint64(dy) + 1) * uint64(z.sh) / (2 * uint64(z.dh))
		d := dst.PixOffset(dp.X+adr.Min.X, dp.Y+int(dy))
		for dx := int32(adr.Min.X); dx < int32(adr.Max.X); dx++ {
			sx := (2*uint64(dx) + 1) * uint64(z.sw) / (2 * uint64(z.dw))
			pi := src.PixOffset(sp.X+int(sx), sp.Y+int(sy))
			pr := uint32(src.Pix[pi+0]) * 0x101
			pg := uint32(src.Pix[pi+1]) * 0x101
			pb := uint32(src.Pix[pi+2]) * 0x101
			pa := uint32(src.Pix[pi+3]) * 0x101
			dst.Pix[d+0] = uint8(uint32(pr) >> 8)
			dst.Pix[d+1] = uint8(uint32(pg) >> 8)
			dst.Pix[d+2] = uint8(uint32(pb) >> 8)
			dst.Pix[d+3] = uint8(uint32(pa) >> 8)
			d += 4
		}
	}
}

func (z *nnScaler) scale_RGBA_Uniform(dst *image.RGBA, dp image.Point, adr image.Rectangle, src *image.Uniform, sp image.Point) {
	for dy := int32(adr.Min.Y); dy < int32(adr.Max.Y); dy++ {
		sy := (2*uint64(dy) + 1) * uint64(z.sh) / (2 * uint64(z.dh))
		d := dst.PixOffset(dp.X+adr.Min.X, dp.Y+int(dy))
		for dx := int32(adr.Min.X); dx < int32(adr.Max.X); dx++ {
			sx := (2*uint64(dx) + 1) * uint64(z.sw) / (2 * uint64(z.dw))
			pr, pg, pb, pa := src.At(sp.X+int(sx), sp.Y+int(sy)).RGBA()
			dst.Pix[d+0] = uint8(uint32(pr) >> 8)
			dst.Pix[d+1] = uint8(uint32(pg) >> 8)
			dst.Pix[d+2] = uint8(uint32(pb) >> 8)
			dst.Pix[d+3] = uint8(uint32(pa) >> 8)
			d += 4
		}
	}
}

func (z *nnScaler) scale_RGBA_YCbCr(dst *image.RGBA, dp image.Point, adr image.Rectangle, src *image.YCbCr, sp image.Point) {
	for dy := int32(adr.Min.Y); dy < int32(adr.Max.Y); dy++ {
		sy := (2*uint64(dy) + 1) * uint64(z.sh) / (2 * uint64(z.dh))
		d := dst.PixOffset(dp.X+adr.Min.X, dp.Y+int(dy))
		for dx := int32(adr.Min.X); dx < int32(adr.Max.X); dx++ {
			sx := (2*uint64(dx) + 1) * uint64(z.sw) / (2 * uint64(z.dw))
			pr, pg, pb, pa := src.At(sp.X+int(sx), sp.Y+int(sy)).RGBA()
			dst.Pix[d+0] = uint8(uint32(pr) >> 8)
			dst.Pix[d+1] = uint8(uint32(pg) >> 8)
			dst.Pix[d+2] = uint8(uint32(pb) >> 8)
			dst.Pix[d+3] = uint8(uint32(pa) >> 8)
			d += 4
		}
	}
}

func (z *nnScaler) scale_RGBA_Image(dst *image.RGBA, dp image.Point, adr image.Rectangle, src image.Image, sp image.Point) {
	for dy := int32(adr.Min.Y); dy < int32(adr.Max.Y); dy++ {
		sy := (2*uint64(dy) + 1) * uint64(z.sh) / (2 * uint64(z.dh))
		d := dst.PixOffset(dp.X+adr.Min.X, dp.Y+int(dy))
		for dx := int32(adr.Min.X); dx < int32(adr.Max.X); dx++ {
			sx := (2*uint64(dx) + 1) * uint64(z.sw) / (2 * uint64(z.dw))
			pr, pg, pb, pa := src.At(sp.X+int(sx), sp.Y+int(sy)).RGBA()
			dst.Pix[d+0] = uint8(uint32(pr) >> 8)
			dst.Pix[d+1] = uint8(uint32(pg) >> 8)
			dst.Pix[d+2] = uint8(uint32(pb) >> 8)
			dst.Pix[d+3] = uint8(uint32(pa) >> 8)
			d += 4
		}
	}
}

func (z *nnScaler) scale_Image_Image(dst Image, dp image.Point, adr image.Rectangle, src image.Image, sp image.Point) {
	dstColorRGBA64 := &color.RGBA64{}
	dstColor := color.Color(dstColorRGBA64)
	for dy := int32(adr.Min.Y); dy < int32(adr.Max.Y); dy++ {
		sy := (2*uint64(dy) + 1) * uint64(z.sh) / (2 * uint64(z.dh))
		for dx := int32(adr.Min.X); dx < int32(adr.Max.X); dx++ {
			sx := (2*uint64(dx) + 1) * uint64(z.sw) / (2 * uint64(z.dw))
			pr, pg, pb, pa := src.At(sp.X+int(sx), sp.Y+int(sy)).RGBA()
			dstColorRGBA64.R = uint16(pr)
			dstColorRGBA64.G = uint16(pg)
			dstColorRGBA64.B = uint16(pb)
			dstColorRGBA64.A = uint16(pa)
			dst.Set(dp.X+int(dx), dp.Y+int(dy), dstColor)
		}
	}
}

func (z *ablScaler) Scale(dst Image, dp image.Point, src image.Image, sp image.Point) {
	if z.dw <= 0 || z.dh <= 0 || z.sw <= 0 || z.sh <= 0 {
		return
	}
	// adr is the affected destination pixels, relative to dp.
	adr := dst.Bounds().Sub(dp).Intersect(image.Rectangle{Max: image.Point{int(z.dw), int(z.dh)}})
	if adr.Empty() {
		return
	}
	// sr is the source pixels. If it extends beyond the src bounds,
	// we cannot use the type-specific fast paths, as they access
	// the Pix fields directly without bounds checking.
	if sr := (image.Rectangle{sp, sp.Add(image.Point{int(z.sw), int(z.sh)})}); !sr.In(src.Bounds()) {
		z.scale_Image_Image(dst, dp, adr, src, sp)
	} else {
		switch dst := dst.(type) {
		case *image.RGBA:
			switch src := src.(type) {
			case *image.Gray:
				z.scale_RGBA_Gray(dst, dp, adr, src, sp)
			case *image.NRGBA:
				z.scale_RGBA_NRGBA(dst, dp, adr, src, sp)
			case *image.RGBA:
				z.scale_RGBA_RGBA(dst, dp, adr, src, sp)
			case *image.Uniform:
				z.scale_RGBA_Uniform(dst, dp, adr, src, sp)
			case *image.YCbCr:
				z.scale_RGBA_YCbCr(dst, dp, adr, src, sp)
			default:
				z.scale_RGBA_Image(dst, dp, adr, src, sp)
			}
		default:
			switch src := src.(type) {
			default:
				z.scale_Image_Image(dst, dp, adr, src, sp)
			}
		}
	}
}

func (z *ablScaler) scale_RGBA_Gray(dst *image.RGBA, dp image.Point, adr image.Rectangle, src *image.Gray, sp image.Point) {
	yscale := float64(z.sh) / float64(z.dh)
	xscale := float64(z.sw) / float64(z.dw)
	for dy := int32(adr.Min.Y); dy < int32(adr.Max.Y); dy++ {
		sy := (float64(dy)+0.5)*yscale - 0.5
		sy0 := int32(sy)
		yFrac0 := sy - float64(sy0)
		yFrac1 := 1 - yFrac0
		sy1 := sy0 + 1
		if sy < 0 {
			sy0, sy1 = 0, 0
			yFrac0, yFrac1 = 0, 1
		} else if sy1 >= z.sh {
			sy1 = sy0
			yFrac0, yFrac1 = 1, 0
		}
		d := dst.PixOffset(dp.X+adr.Min.X, dp.Y+int(dy))
		for dx := int32(adr.Min.X); dx < int32(adr.Max.X); dx++ {
			sx := (float64(dx)+0.5)*xscale - 0.5
			sx0 := int32(sx)
			xFrac0 := sx - float64(sx0)
			xFrac1 := 1 - xFrac0
			sx1 := sx0 + 1
			if sx < 0 {
				sx0, sx1 = 0, 0
				xFrac0, xFrac1 = 0, 1
			} else if sx1 >= z.sw {
				sx1 = sx0
				xFrac0, xFrac1 = 1, 0
			}
			s00ru, s00gu, s00bu, s00au := src.At(sp.X+int(sx0), sp.Y+int(sy0)).RGBA()
			s00r := float64(s00ru)
			s00g := float64(s00gu)
			s00b := float64(s00bu)
			s00a := float64(s00au)
			s10ru, s10gu, s10bu, s10au := src.At(sp.X+int(sx1), sp.Y+int(sy0)).RGBA()
			s10r := float64(s10ru)
			s10g := float64(s10gu)
			s10b := float64(s10bu)
			s10a := float64(s10au)
			s10r = xFrac1*s00r + xFrac0*s10r
			s10g = xFrac1*s00g + xFrac0*s10g
			s10b = xFrac1*s00b + xFrac0*s10b
			s10a = xFrac1*s00a + xFrac0*s10a
			s01ru, s01gu, s01bu, s01au := src.At(sp.X+int(sx0), sp.Y+int(sy1)).RGBA()
			s01r := float64(s01ru)
			s01g := float64(s01gu)
			s01b := float64(s01bu)
			s01a := float64(s01au)
			s11ru, s11gu, s11bu, s11au := src.At(sp.X+int(sx1), sp.Y+int(sy1)).RGBA()
			s11r := float64(s11ru)
			s11g := float64(s11gu)
			s11b := float64(s11bu)
			s11a := float64(s11au)
			s11r = xFrac1*s01r + xFrac0*s11r
			s11g = xFrac1*s01g + xFrac0*s11g
			s11b = xFrac1*s01b + xFrac0*s11b
			s11a = xFrac1*s01a + xFrac0*s11a
			s11r = yFrac1*s10r + yFrac0*s11r
			s11g = yFrac1*s10g + yFrac0*s11g
			s11b = yFrac1*s10b + yFrac0*s11b
			s11a = yFrac1*s10a + yFrac0*s11a
			dst.Pix[d+0] = uint8(uint32(s11r) >> 8)
			dst.Pix[d+1] = uint8(uint32(s11g) >> 8)
			dst.Pix[d+2] = uint8(uint32(s11b) >> 8)
			dst.Pix[d+3] = uint8(uint32(s11a) >> 8)
			d += 4
		}
	}
}

func (z *ablScaler) scale_RGBA_NRGBA(dst *image.RGBA, dp image.Point, adr image.Rectangle, src *image.NRGBA, sp image.Point) {
	yscale := float64(z.sh) / float64(z.dh)
	xscale := float64(z.sw) / float64(z.dw)
	for dy := int32(adr.Min.Y); dy < int32(adr.Max.Y); dy++ {
		sy := (float64(dy)+0.5)*yscale - 0.5
		sy0 := int32(sy)
		yFrac0 := sy - float64(sy0)
		yFrac1 := 1 - yFrac0
		sy1 := sy0 + 1
		if sy < 0 {
			sy0, sy1 = 0, 0
			yFrac0, yFrac1 = 0, 1
		} else if sy1 >= z.sh {
			sy1 = sy0
			yFrac0, yFrac1 = 1, 0
		}
		d := dst.PixOffset(dp.X+adr.Min.X, dp.Y+int(dy))
		for dx := int32(adr.Min.X); dx < int32(adr.Max.X); dx++ {
			sx := (float64(dx)+0.5)*xscale - 0.5
			sx0 := int32(sx)
			xFrac0 := sx - float64(sx0)
			xFrac1 := 1 - xFrac0
			sx1 := sx0 + 1
			if sx < 0 {
				sx0, sx1 = 0, 0
				xFrac0, xFrac1 = 0, 1
			} else if sx1 >= z.sw {
				sx1 = sx0
				xFrac0, xFrac1 = 1, 0
			}
			s00ru, s00gu, s00bu, s00au := src.At(sp.X+int(sx0), sp.Y+int(sy0)).RGBA()
			s00r := float64(s00ru)
			s00g := float64(s00gu)
			s00b := float64(s00bu)
			s00a := float64(s00au)
			s10ru, s10gu, s10bu, s10au := src.At(sp.X+int(sx1), sp.Y+int(sy0)).RGBA()
			s10r := float64(s10ru)
			s10g := float64(s10gu)
			s10b := float64(s10bu)
			s10a := float64(s10au)
			s10r = xFrac1*s00r + xFrac0*s10r
			s10g = xFrac1*s00g + xFrac0*s10g
			s10b = xFrac1*s00b + xFrac0*s10b
			s10a = xFrac1*s00a + xFrac0*s10a
			s01ru, s01gu, s01bu, s01au := src.At(sp.X+int(sx0), sp.Y+int(sy1)).RGBA()
			s01r := float64(s01ru)
			s01g := float64(s01gu)
			s01b := float64(s01bu)
			s01a := float64(s01au)
			s11ru, s11gu, s11bu, s11au := src.At(sp.X+int(sx1), sp.Y+int(sy1)).RGBA()
			s11r := float64(s11ru)
			s11g := float64(s11gu)
			s11b := float64(s11bu)
			s11a := float64(s11au)
			s11r = xFrac1*s01r + xFrac0*s11r
			s11g = xFrac1*s01g + xFrac0*s11g
			s11b = xFrac1*s01b + xFrac0*s11b
			s11a = xFrac1*s01a + xFrac0*s11a
			s11r = yFrac1*s10r + yFrac0*s11r
			s11g = yFrac1*s10g + yFrac0*s11g
			s11b = yFrac1*s10b + yFrac0*s11b
			s11a = yFrac1*s10a + yFrac0*s11a
			dst.Pix[d+0] = uint8(uint32(s11r) >> 8)
			dst.Pix[d+1] = uint8(uint32(s11g) >> 8)
			dst.Pix[d+2] = uint8(uint32(s11b) >> 8)
			dst.Pix[d+3] = uint8(uint32(s11a) >> 8)
			d += 4
		}
	}
}

func (z *ablScaler) scale_RGBA_RGBA(dst *image.RGBA, dp image.Point, adr image.Rectangle, src *image.RGBA, sp image.Point) {
	yscale := float64(z.sh) / float64(z.dh)
	xscale := float64(z.sw) / float64(z.dw)
	for dy := int32(adr.Min.Y); dy < int32(adr.Max.Y); dy++ {
		sy := (float64(dy)+0.5)*yscale - 0.5
		sy0 := int32(sy)
		yFrac0 := sy - float64(sy0)
		yFrac1 := 1 - yFrac0
		sy1 := sy0 + 1
		if sy < 0 {
			sy0, sy1 = 0, 0
			yFrac0, yFrac1 = 0, 1
		} else if sy1 >= z.sh {
			sy1 = sy0
			yFrac0, yFrac1 = 1, 0
		}
		d := dst.PixOffset(dp.X+adr.Min.X, dp.Y+int(dy))
		for dx := int32(adr.Min.X); dx < int32(adr.Max.X); dx++ {
			sx := (float64(dx)+0.5)*xscale - 0.5
			sx0 := int32(sx)
			xFrac0 := sx - float64(sx0)
			xFrac1 := 1 - xFrac0
			sx1 := sx0 + 1
			if sx < 0 {
				sx0, sx1 = 0, 0
				xFrac0, xFrac1 = 0, 1
			} else if sx1 >= z.sw {
				sx1 = sx0
				xFrac0, xFrac1 = 1, 0
			}
			s00i := src.PixOffset(sp.X+int(sx0), sp.Y+int(sy0))
			s00ru := uint32(src.Pix[s00i+0]) * 0x101
			s00gu := uint32(src.Pix[s00i+1]) * 0x101
			s00bu := uint32(src.Pix[s00i+2]) * 0x101
			s00au := uint32(src.Pix[s00i+3]) * 0x101
			s00r := float64(s00ru)
			s00g := float64(s00gu)
			s00b := float64(s00bu)
			s00a := float64(s00au)
			s10i := src.PixOffset(sp.X+int(sx1), sp.Y+int(sy0))
			s10ru := uint32(src.Pix[s10i+0]) * 0x101
			s10gu := uint32(src.Pix[s10i+1]) * 0x101
			s10bu := uint32(src.Pix[s10i+2]) * 0x101
			s10au := uint32(src.Pix[s10i+3]) * 0x101
			s10r := float64(s10ru)
			s10g := float64(s10gu)
			s10b := float64(s10bu)
			s10a := float64(s10au)
			s10r = xFrac1*s00r + xFrac0*s10r
			s10g = xFrac1*s00g + xFrac0*s10g
			s10b = xFrac1*s00b + xFrac0*s10b
			s10a = xFrac1*s00a + xFrac0*s10a
			s01i := src.PixOffset(sp.X+int(sx0), sp.Y+int(sy1))
			s01ru := uint32(src.Pix[s01i+0]) * 0x101
			s01gu := uint32(src.Pix[s01i+1]) * 0x101
			s01bu := uint32(src.Pix[s01i+2]) * 0x101
			s01au := uint32(src.Pix[s01i+3]) * 0x101
			s01r := float64(s01ru)
			s01g := float64(s01gu)
			s01b := float64(s01bu)
			s01a := float64(s01au)
			s11i := src.PixOffset(sp.X+int(sx1), sp.Y+int(sy1))
			s11ru := uint32(src.Pix[s11i+0]) * 0x101
			s11gu := uint32(src.Pix[s11i+1]) * 0x101
			s11bu := uint32(src.Pix[s11i+2]) * 0x101
			s11au := uint32(src.Pix[s11i+3]) * 0x101
			s11r := float64(s11ru)
			s11g := float64(s11gu)
			s11b := float64(s11bu)
			s11a := float64(s11au)
			s11r = xFrac1*s01r + xFrac0*s11r
			s11g = xFrac1*s01g + xFrac0*s11g
			s11b = xFrac1*s01b + xFrac0*s11b
			s11a = xFrac1*s01a + xFrac0*s11a
			s11r = yFrac1*s10r + yFrac0*s11r
			s11g = yFrac1*s10g + yFrac0*s11g
			s11b = yFrac1*s10b + yFrac0*s11b
			s11a = yFrac1*s10a + yFrac0*s11a
			dst.Pix[d+0] = uint8(uint32(s11r) >> 8)
			dst.Pix[d+1] = uint8(uint32(s11g) >> 8)
			dst.Pix[d+2] = uint8(uint32(s11b) >> 8)
			dst.Pix[d+3] = uint8(uint32(s11a) >> 8)
			d += 4
		}
	}
}

func (z *ablScaler) scale_RGBA_Uniform(dst *image.RGBA, dp image.Point, adr image.Rectangle, src *image.Uniform, sp image.Point) {
	yscale := float64(z.sh) / float64(z.dh)
	xscale := float64(z.sw) / float64(z.dw)
	for dy := int32(adr.Min.Y); dy < int32(adr.Max.Y); dy++ {
		sy := (float64(dy)+0.5)*yscale - 0.5
		sy0 := int32(sy)
		yFrac0 := sy - float64(sy0)
		yFrac1 := 1 - yFrac0
		sy1 := sy0 + 1
		if sy < 0 {
			sy0, sy1 = 0, 0
			yFrac0, yFrac1 = 0, 1
		} else if sy1 >= z.sh {
			sy1 = sy0
			yFrac0, yFrac1 = 1, 0
		}
		d := dst.PixOffset(dp.X+adr.Min.X, dp.Y+int(dy))
		for dx := int32(adr.Min.X); dx < int32(adr.Max.X); dx++ {
			sx := (float64(dx)+0.5)*xscale - 0.5
			sx0 := int32(sx)
			xFrac0 := sx - float64(sx0)
			xFrac1 := 1 - xFrac0
			sx1 := sx0 + 1
			if sx < 0 {
				sx0, sx1 = 0, 0
				xFrac0, xFrac1 = 0, 1
			} else if sx1 >= z.sw {
				sx1 = sx0
				xFrac0, xFrac1 = 1, 0
			}
			s00ru, s00gu, s00bu, s00au := src.At(sp.X+int(sx0), sp.Y+int(sy0)).RGBA()
			s00r := float64(s00ru)
			s00g := float64(s00gu)
			s00b := float64(s00bu)
			s00a := float64(s00au)
			s10ru, s10gu, s10bu, s10au := src.At(sp.X+int(sx1), sp.Y+int(sy0)).RGBA()
			s10r := float64(s10ru)
			s10g := float64(s10gu)
			s10b := float64(s10bu)
			s10a := float64(s10au)
			s10r = xFrac1*s00r + xFrac0*s10r
			s10g = xFrac1*s00g + xFrac0*s10g
			s10b = xFrac1*s00b + xFrac0*s10b
			s10a = xFrac1*s00a + xFrac0*s10a
			s01ru, s01gu, s01bu, s01au := src.At(sp.X+int(sx0), sp.Y+int(sy1)).RGBA()
			s01r := float64(s01ru)
			s01g := float64(s01gu)
			s01b := float64(s01bu)
			s01a := float64(s01au)
			s11ru, s11gu, s11bu, s11au := src.At(sp.X+int(sx1), sp.Y+int(sy1)).RGBA()
			s11r := float64(s11ru)
			s11g := float64(s11gu)
			s11b := float64(s11bu)
			s11a := float64(s11au)
			s11r = xFrac1*s01r + xFrac0*s11r
			s11g = xFrac1*s01g + xFrac0*s11g
			s11b = xFrac1*s01b + xFrac0*s11b
			s11a = xFrac1*s01a + xFrac0*s11a
			s11r = yFrac1*s10r + yFrac0*s11r
			s11g = yFrac1*s10g + yFrac0*s11g
			s11b = yFrac1*s10b + yFrac0*s11b
			s11a = yFrac1*s10a + yFrac0*s11a
			dst.Pix[d+0] = uint8(uint32(s11r) >> 8)
			dst.Pix[d+1] = uint8(uint32(s11g) >> 8)
			dst.Pix[d+2] = uint8(uint32(s11b) >> 8)
			dst.Pix[d+3] = uint8(uint32(s11a) >> 8)
			d += 4
		}
	}
}

func (z *ablScaler) scale_RGBA_YCbCr(dst *image.RGBA, dp image.Point, adr image.Rectangle, src *image.YCbCr, sp image.Point) {
	yscale := float64(z.sh) / float64(z.dh)
	xscale := float64(z.sw) / float64(z.dw)
	for dy := int32(adr.Min.Y); dy < int32(adr.Max.Y); dy++ {
		sy := (float64(dy)+0.5)*yscale - 0.5
		sy0 := int32(sy)
		yFrac0 := sy - float64(sy0)
		yFrac1 := 1 - yFrac0
		sy1 := sy0 + 1
		if sy < 0 {
			sy0, sy1 = 0, 0
			yFrac0, yFrac1 = 0, 1
		} else if sy1 >= z.sh {
			sy1 = sy0
			yFrac0, yFrac1 = 1, 0
		}
		d := dst.PixOffset(dp.X+adr.Min.X, dp.Y+int(dy))
		for dx := int32(adr.Min.X); dx < int32(adr.Max.X); dx++ {
			sx := (float64(dx)+0.5)*xscale - 0.5
			sx0 := int32(sx)
			xFrac0 := sx - float64(sx0)
			xFrac1 := 1 - xFrac0
			sx1 := sx0 + 1
			if sx < 0 {
				sx0, sx1 = 0, 0
				xFrac0, xFrac1 = 0, 1
			} else if sx1 >= z.sw {
				sx1 = sx0
				xFrac0, xFrac1 = 1, 0
			}
			s00ru, s00gu, s00bu, s00au := src.At(sp.X+int(sx0), sp.Y+int(sy0)).RGBA()
			s00r := float64(s00ru)
			s00g := float64(s00gu)
			s00b := float64(s00bu)
			s00a := float64(s00au)
			s10ru, s10gu, s10bu, s10au := src.At(sp.X+int(sx1), sp.Y+int(sy0)).RGBA()
			s10r := float64(s10ru)
			s10g := float64(s10gu)
			s10b := float64(s10bu)
			s10a := float64(s10au)
			s10r = xFrac1*s00r + xFrac0*s10r
			s10g = xFrac1*s00g + xFrac0*s10g
			s10b = xFrac1*s00b + xFrac0*s10b
			s10a = xFrac1*s00a + xFrac0*s10a
			s01ru, s01gu, s01bu, s01au := src.At(sp.X+int(sx0), sp.Y+int(sy1)).RGBA()
			s01r := float64(s01ru)
			s01g := float64(s01gu)
			s01b := float64(s01bu)
			s01a := float64(s01au)
			s11ru, s11gu, s11bu, s11au := src.At(sp.X+int(sx1), sp.Y+int(sy1)).RGBA()
			s11r := float64(s11ru)
			s11g := float64(s11gu)
			s11b := float64(s11bu)
			s11a := float64(s11au)
			s11r = xFrac1*s01r + xFrac0*s11r
			s11g = xFrac1*s01g + xFrac0*s11g
			s11b = xFrac1*s01b + xFrac0*s11b
			s11a = xFrac1*s01a + xFrac0*s11a
			s11r = yFrac1*s10r + yFrac0*s11r
			s11g = yFrac1*s10g + yFrac0*s11g
			s11b = yFrac1*s10b + yFrac0*s11b
			s11a = yFrac1*s10a + yFrac0*s11a
			dst.Pix[d+0] = uint8(uint32(s11r) >> 8)
			dst.Pix[d+1] = uint8(uint32(s11g) >> 8)
			dst.Pix[d+2] = uint8(uint32(s11b) >> 8)
			dst.Pix[d+3] = uint8(uint32(s11a) >> 8)
			d += 4
		}
	}
}

func (z *ablScaler) scale_RGBA_Image(dst *image.RGBA, dp image.Point, adr image.Rectangle, src image.Image, sp image.Point) {
	yscale := float64(z.sh) / float64(z.dh)
	xscale := float64(z.sw) / float64(z.dw)
	for dy := int32(adr.Min.Y); dy < int32(adr.Max.Y); dy++ {
		sy := (float64(dy)+0.5)*yscale - 0.5
		sy0 := int32(sy)
		yFrac0 := sy - float64(sy0)
		yFrac1 := 1 - yFrac0
		sy1 := sy0 + 1
		if sy < 0 {
			sy0, sy1 = 0, 0
			yFrac0, yFrac1 = 0, 1
		} else if sy1 >= z.sh {
			sy1 = sy0
			yFrac0, yFrac1 = 1, 0
		}
		d := dst.PixOffset(dp.X+adr.Min.X, dp.Y+int(dy))
		for dx := int32(adr.Min.X); dx < int32(adr.Max.X); dx++ {
			sx := (float64(dx)+0.5)*xscale - 0.5
			sx0 := int32(sx)
			xFrac0 := sx - float64(sx0)
			xFrac1 := 1 - xFrac0
			sx1 := sx0 + 1
			if sx < 0 {
				sx0, sx1 = 0, 0
				xFrac0, xFrac1 = 0, 1
			} else if sx1 >= z.sw {
				sx1 = sx0
				xFrac0, xFrac1 = 1, 0
			}
			s00ru, s00gu, s00bu, s00au := src.At(sp.X+int(sx0), sp.Y+int(sy0)).RGBA()
			s00r := float64(s00ru)
			s00g := float64(s00gu)
			s00b := float64(s00bu)
			s00a := float64(s00au)
			s10ru, s10gu, s10bu, s10au := src.At(sp.X+int(sx1), sp.Y+int(sy0)).RGBA()
			s10r := float64(s10ru)
			s10g := float64(s10gu)
			s10b := float64(s10bu)
			s10a := float64(s10au)
			s10r = xFrac1*s00r + xFrac0*s10r
			s10g = xFrac1*s00g + xFrac0*s10g
			s10b = xFrac1*s00b + xFrac0*s10b
			s10a = xFrac1*s00a + xFrac0*s10a
			s01ru, s01gu, s01bu, s01au := src.At(sp.X+int(sx0), sp.Y+int(sy1)).RGBA()
			s01r := float64(s01ru)
			s01g := float64(s01gu)
			s01b := float64(s01bu)
			s01a := float64(s01au)
			s11ru, s11gu, s11bu, s11au := src.At(sp.X+int(sx1), sp.Y+int(sy1)).RGBA()
			s11r := float64(s11ru)
			s11g := float64(s11gu)
			s11b := float64(s11bu)
			s11a := float64(s11au)
			s11r = xFrac1*s01r + xFrac0*s11r
			s11g = xFrac1*s01g + xFrac0*s11g
			s11b = xFrac1*s01b + xFrac0*s11b
			s11a = xFrac1*s01a + xFrac0*s11a
			s11r = yFrac1*s10r + yFrac0*s11r
			s11g = yFrac1*s10g + yFrac0*s11g
			s11b = yFrac1*s10b + yFrac0*s11b
			s11a = yFrac1*s10a + yFrac0*s11a
			dst.Pix[d+0] = uint8(uint32(s11r) >> 8)
			dst.Pix[d+1] = uint8(uint32(s11g) >> 8)
			dst.Pix[d+2] = uint8(uint32(s11b) >> 8)
			dst.Pix[d+3] = uint8(uint32(s11a) >> 8)
			d += 4
		}
	}
}

func (z *ablScaler) scale_Image_Image(dst Image, dp image.Point, adr image.Rectangle, src image.Image, sp image.Point) {
	yscale := float64(z.sh) / float64(z.dh)
	xscale := float64(z.sw) / float64(z.dw)
	dstColorRGBA64 := &color.RGBA64{}
	dstColor := color.Color(dstColorRGBA64)
	for dy := int32(adr.Min.Y); dy < int32(adr.Max.Y); dy++ {
		sy := (float64(dy)+0.5)*yscale - 0.5
		sy0 := int32(sy)
		yFrac0 := sy - float64(sy0)
		yFrac1 := 1 - yFrac0
		sy1 := sy0 + 1
		if sy < 0 {
			sy0, sy1 = 0, 0
			yFrac0, yFrac1 = 0, 1
		} else if sy1 >= z.sh {
			sy1 = sy0
			yFrac0, yFrac1 = 1, 0
		}
		for dx := int32(adr.Min.X); dx < int32(adr.Max.X); dx++ {
			sx := (float64(dx)+0.5)*xscale - 0.5
			sx0 := int32(sx)
			xFrac0 := sx - float64(sx0)
			xFrac1 := 1 - xFrac0
			sx1 := sx0 + 1
			if sx < 0 {
				sx0, sx1 = 0, 0
				xFrac0, xFrac1 = 0, 1
			} else if sx1 >= z.sw {
				sx1 = sx0
				xFrac0, xFrac1 = 1, 0
			}
			s00ru, s00gu, s00bu, s00au := src.At(sp.X+int(sx0), sp.Y+int(sy0)).RGBA()
			s00r := float64(s00ru)
			s00g := float64(s00gu)
			s00b := float64(s00bu)
			s00a := float64(s00au)
			s10ru, s10gu, s10bu, s10au := src.At(sp.X+int(sx1), sp.Y+int(sy0)).RGBA()
			s10r := float64(s10ru)
			s10g := float64(s10gu)
			s10b := float64(s10bu)
			s10a := float64(s10au)
			s10r = xFrac1*s00r + xFrac0*s10r
			s10g = xFrac1*s00g + xFrac0*s10g
			s10b = xFrac1*s00b + xFrac0*s10b
			s10a = xFrac1*s00a + xFrac0*s10a
			s01ru, s01gu, s01bu, s01au := src.At(sp.X+int(sx0), sp.Y+int(sy1)).RGBA()
			s01r := float64(s01ru)
			s01g := float64(s01gu)
			s01b := float64(s01bu)
			s01a := float64(s01au)
			s11ru, s11gu, s11bu, s11au := src.At(sp.X+int(sx1), sp.Y+int(sy1)).RGBA()
			s11r := float64(s11ru)
			s11g := float64(s11gu)
			s11b := float64(s11bu)
			s11a := float64(s11au)
			s11r = xFrac1*s01r + xFrac0*s11r
			s11g = xFrac1*s01g + xFrac0*s11g
			s11b = xFrac1*s01b + xFrac0*s11b
			s11a = xFrac1*s01a + xFrac0*s11a
			s11r = yFrac1*s10r + yFrac0*s11r
			s11g = yFrac1*s10g + yFrac0*s11g
			s11b = yFrac1*s10b + yFrac0*s11b
			s11a = yFrac1*s10a + yFrac0*s11a
			dstColorRGBA64.R = uint16(s11r)
			dstColorRGBA64.G = uint16(s11g)
			dstColorRGBA64.B = uint16(s11b)
			dstColorRGBA64.A = uint16(s11a)
			dst.Set(dp.X+int(dx), dp.Y+int(dy), dstColor)
		}
	}
}

func (z *kernelScaler) Scale(dst Image, dp image.Point, src image.Image, sp image.Point) {
	if z.dw <= 0 || z.dh <= 0 || z.sw <= 0 || z.sh <= 0 {
		return
	}
	// adr is the affected destination pixels, relative to dp.
	adr := dst.Bounds().Sub(dp).Intersect(image.Rectangle{Max: image.Point{int(z.dw), int(z.dh)}})
	if adr.Empty() {
		return
	}
	// Create a temporary buffer:
	// scaleX distributes the source image's columns over the temporary image.
	// scaleY distributes the temporary image's rows over the destination image.
	// TODO: is it worth having a sync.Pool for this temporary buffer?
	tmp := make([][4]float64, z.dw*z.sh)

	// sr is the source pixels. If it extends beyond the src bounds,
	// we cannot use the type-specific fast paths, as they access
	// the Pix fields directly without bounds checking.
	if sr := (image.Rectangle{sp, sp.Add(image.Point{int(z.sw), int(z.sh)})}); !sr.In(src.Bounds()) {
		z.scaleX_Image(tmp, src, sp)
	} else {
		switch src := src.(type) {
		case *image.Gray:
			z.scaleX_Gray(tmp, src, sp)
		case *image.NRGBA:
			z.scaleX_NRGBA(tmp, src, sp)
		case *image.RGBA:
			z.scaleX_RGBA(tmp, src, sp)
		case *image.Uniform:
			z.scaleX_Uniform(tmp, src, sp)
		case *image.YCbCr:
			z.scaleX_YCbCr(tmp, src, sp)
		default:
			z.scaleX_Image(tmp, src, sp)
		}
	}

	switch dst := dst.(type) {
	case *image.RGBA:
		z.scaleY_RGBA(dst, dp, adr, tmp)
	default:
		z.scaleY_Image(dst, dp, adr, tmp)
	}
}

func (z *kernelScaler) scaleX_Gray(tmp [][4]float64, src *image.Gray, sp image.Point) {
	t := 0
	for y := int32(0); y < z.sh; y++ {
		for _, s := range z.horizontal.sources {
			var pr, pg, pb, pa float64
			for _, c := range z.horizontal.contribs[s.i:s.j] {
				pru, pgu, pbu, pau := src.At(sp.X+int(c.coord), sp.Y+int(y)).RGBA()
				pr += float64(pru) * c.weight
				pg += float64(pgu) * c.weight
				pb += float64(pbu) * c.weight
				pa += float64(pau) * c.weight
			}
			tmp[t] = [4]float64{
				pr * s.invTotalWeightFFFF,
				pg * s.invTotalWeightFFFF,
				pb * s.invTotalWeightFFFF,
				pa * s.invTotalWeightFFFF,
			}
			t++
		}
	}
}

func (z *kernelScaler) scaleX_NRGBA(tmp [][4]float64, src *image.NRGBA, sp image.Point) {
	t := 0
	for y := int32(0); y < z.sh; y++ {
		for _, s := range z.horizontal.sources {
			var pr, pg, pb, pa float64
			for _, c := range z.horizontal.contribs[s.i:s.j] {
				pru, pgu, pbu, pau := src.At(sp.X+int(c.coord), sp.Y+int(y)).RGBA()
				pr += float64(pru) * c.weight
				pg += float64(pgu) * c.weight
				pb += float64(pbu) * c.weight
				pa += float64(pau) * c.weight
			}
			tmp[t] = [4]float64{
				pr * s.invTotalWeightFFFF,
				pg * s.invTotalWeightFFFF,
				pb * s.invTotalWeightFFFF,
				pa * s.invTotalWeightFFFF,
			}
			t++
		}
	}
}

func (z *kernelScaler) scaleX_RGBA(tmp [][4]float64, src *image.RGBA, sp image.Point) {
	t := 0
	for y := int32(0); y < z.sh; y++ {
		for _, s := range z.horizontal.sources {
			var pr, pg, pb, pa float64
			for _, c := range z.horizontal.contribs[s.i:s.j] {
				pi := src.PixOffset(sp.X+int(c.coord), sp.Y+int(y))
				pru := uint32(src.Pix[pi+0]) * 0x101
				pgu := uint32(src.Pix[pi+1]) * 0x101
				pbu := uint32(src.Pix[pi+2]) * 0x101
				pau := uint32(src.Pix[pi+3]) * 0x101
				pr += float64(pru) * c.weight
				pg += float64(pgu) * c.weight
				pb += float64(pbu) * c.weight
				pa += float64(pau) * c.weight
			}
			tmp[t] = [4]float64{
				pr * s.invTotalWeightFFFF,
				pg * s.invTotalWeightFFFF,
				pb * s.invTotalWeightFFFF,
				pa * s.invTotalWeightFFFF,
			}
			t++
		}
	}
}

func (z *kernelScaler) scaleX_Uniform(tmp [][4]float64, src *image.Uniform, sp image.Point) {
	t := 0
	for y := int32(0); y < z.sh; y++ {
		for _, s := range z.horizontal.sources {
			var pr, pg, pb, pa float64
			for _, c := range z.horizontal.contribs[s.i:s.j] {
				pru, pgu, pbu, pau := src.At(sp.X+int(c.coord), sp.Y+int(y)).RGBA()
				pr += float64(pru) * c.weight
				pg += float64(pgu) * c.weight
				pb += float64(pbu) * c.weight
				pa += float64(pau) * c.weight
			}
			tmp[t] = [4]float64{
				pr * s.invTotalWeightFFFF,
				pg * s.invTotalWeightFFFF,
				pb * s.invTotalWeightFFFF,
				pa * s.invTotalWeightFFFF,
			}
			t++
		}
	}
}

func (z *kernelScaler) scaleX_YCbCr(tmp [][4]float64, src *image.YCbCr, sp image.Point) {
	t := 0
	for y := int32(0); y < z.sh; y++ {
		for _, s := range z.horizontal.sources {
			var pr, pg, pb, pa float64
			for _, c := range z.horizontal.contribs[s.i:s.j] {
				pru, pgu, pbu, pau := src.At(sp.X+int(c.coord), sp.Y+int(y)).RGBA()
				pr += float64(pru) * c.weight
				pg += float64(pgu) * c.weight
				pb += float64(pbu) * c.weight
				pa += float64(pau) * c.weight
			}
			tmp[t] = [4]float64{
				pr * s.invTotalWeightFFFF,
				pg * s.invTotalWeightFFFF,
				pb * s.invTotalWeightFFFF,
				pa * s.invTotalWeightFFFF,
			}
			t++
		}
	}
}

func (z *kernelScaler) scaleX_Image(tmp [][4]float64, src image.Image, sp image.Point) {
	t := 0
	for y := int32(0); y < z.sh; y++ {
		for _, s := range z.horizontal.sources {
			var pr, pg, pb, pa float64
			for _, c := range z.horizontal.contribs[s.i:s.j] {
				pru, pgu, pbu, pau := src.At(sp.X+int(c.coord), sp.Y+int(y)).RGBA()
				pr += float64(pru) * c.weight
				pg += float64(pgu) * c.weight
				pb += float64(pbu) * c.weight
				pa += float64(pau) * c.weight
			}
			tmp[t] = [4]float64{
				pr * s.invTotalWeightFFFF,
				pg * s.invTotalWeightFFFF,
				pb * s.invTotalWeightFFFF,
				pa * s.invTotalWeightFFFF,
			}
			t++
		}
	}
}

func (z *kernelScaler) scaleY_RGBA(dst *image.RGBA, dp image.Point, adr image.Rectangle, tmp [][4]float64) {
	for dx := int32(adr.Min.X); dx < int32(adr.Max.X); dx++ {
		d := dst.PixOffset(dp.X+int(dx), dp.Y+adr.Min.Y)
		for _, s := range z.vertical.sources[adr.Min.Y:adr.Max.Y] {
			var pr, pg, pb, pa float64
			for _, c := range z.vertical.contribs[s.i:s.j] {
				p := &tmp[c.coord*z.dw+dx]
				pr += p[0] * c.weight
				pg += p[1] * c.weight
				pb += p[2] * c.weight
				pa += p[3] * c.weight
			}
			dst.Pix[d+0] = uint8(ftou(pr*s.invTotalWeight) >> 8)
			dst.Pix[d+1] = uint8(ftou(pg*s.invTotalWeight) >> 8)
			dst.Pix[d+2] = uint8(ftou(pb*s.invTotalWeight) >> 8)
			dst.Pix[d+3] = uint8(ftou(pa*s.invTotalWeight) >> 8)
			d += dst.Stride
		}
	}
}

func (z *kernelScaler) scaleY_Image(dst Image, dp image.Point, adr image.Rectangle, tmp [][4]float64) {
	dstColorRGBA64 := &color.RGBA64{}
	dstColor := color.Color(dstColorRGBA64)
	for dx := int32(adr.Min.X); dx < int32(adr.Max.X); dx++ {
		for dy, s := range z.vertical.sources[adr.Min.Y:adr.Max.Y] {
			var pr, pg, pb, pa float64
			for _, c := range z.vertical.contribs[s.i:s.j] {
				p := &tmp[c.coord*z.dw+dx]
				pr += p[0] * c.weight
				pg += p[1] * c.weight
				pb += p[2] * c.weight
				pa += p[3] * c.weight
			}
			dstColorRGBA64.R = ftou(pr * s.invTotalWeight)
			dstColorRGBA64.G = ftou(pg * s.invTotalWeight)
			dstColorRGBA64.B = ftou(pb * s.invTotalWeight)
			dstColorRGBA64.A = ftou(pa * s.invTotalWeight)
			dst.Set(dp.X+int(dx), dp.Y+int(adr.Min.Y+dy), dstColor)
		}
	}
}
