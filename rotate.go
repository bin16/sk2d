package sk2d

import (
	"log"
	"math"

	"github.com/fogleman/gg"
)

type Point struct {
	X, Y float64
}

func rotate(x, y, x0, y0, ang float64) (x1, y1 float64) {
	x1 = x*math.Cos(ang) - y*math.Sin(ang)
	y1 = x*math.Sin(ang) + y*math.Cos(ang)
	return
}

func (p *Project) getJointLocation(jName string) (x, y float64) {
	if jName == "" {
		log.Panicf("getJointLocation bad jName: <%s>\n", jName)
	}

	if jName == "root" {
		x = float64(p.Width) / 2
		y = float64(p.Height) / 2
		if p.OffsetX != 0 {
			x = float64(p.Width)/2 + float64(p.OffsetX)
		}
		if s := p.getRootSetter("x"); s != nil {
			x = float64(p.Height)/2 + s.ParseFloat64()
		}
		if p.OffsetY != 0 {
			y = float64(p.Width)/2 + float64(p.OffsetY)
		}
		if s := p.getRootSetter("y"); s != nil {
			y = float64(p.Height)/2 + s.ParseFloat64()
		}
		return
	}

	bName := p.jointCache[jName]
	b := p.getBone(bName)

	if jName == b.Map.From {
		return p.getJointLocation(b.Map.To)
	}

	// abs loc of BoneMap.From/BoneMap.To
	ax, ay := p.getJointLocation(b.Map.To)
	// rel loc of joint (rel to BoneMap.From)
	x1, y1 := intsToFloat64(p.boneGetJointRelLocation(bName, jName))

	dx, dy := rotate(x1, y1, 0, 0, p.getAngle(b.Name))
	x = dx + ax
	y = dy + ay
	return
}

func (p *Project) SaveAs(filename string) {
	// strokeColor := "#0a65d0"
	// fillColor := "f9f9f9"

	strokeColor := "#c8130e"
	fillColor := "#faf1c8"
	if p.Fill != "" {
		fillColor = p.Fill
	}
	if p.Stroke != "" {
		strokeColor = p.Stroke
	}

	ctx := gg.NewContext(p.Width, p.Height)
	ctx.DrawRectangle(0, 0, float64(p.Width), float64(p.Height))
	ctx.SetHexColor(fillColor)
	ctx.SetLineWidth(2)
	ctx.Fill()

	ctx.SetHexColor(strokeColor)
	for _, b := range p.Bones {
		bw, bh := p.getBoneSize(b.Name)
		ang := p.getAngle(b.Name)

		if imgSrc := p.getBoneImageSrc(b.Name); imgSrc != "" {
			ax, ay := p.getJointLocation(b.Map.To)
			x0, y0 := p.boneGetPointRelLocation(b.Name, 0, 0)

			ctx.RotateAbout(ang, ax, ay)
			ctx.DrawImage(p.imageCache[imgSrc], int(ax)+x0, int(ay)+y0)
			ctx.RotateAbout(-ang, ax, ay)
		}

		if !p.debug {
			continue
		}

		for _, d := range [][2]int{
			{0, 0},
			{bw, 0},
			{bw, bh},
			{0, bh},
		} {
			// abs loc of BoneMap.From/BoneMap.To
			ax, ay := p.getJointLocation(b.Map.To)
			x1, y1 := intsToFloat64(p.boneGetPointRelLocation(b.Name, d[0], d[1]))

			dx, dy := rotate(x1, y1, 0, 0, ang)
			x := dx + ax
			y := dy + ay

			ctx.LineTo(x, y)
		}
		ctx.ClosePath()
		if b.Image.Src == "" {
			ctx.SetHexColor(fillColor)
			ctx.FillPreserve()
		}
		ctx.SetHexColor(strokeColor)
		ctx.Stroke()

		for _, j := range b.Joints {
			jx, jy := p.getJointLocation(j.Name)
			ctx.DrawCircle(jx, jy, 6)
			ctx.Stroke()
		}
	}

	ctx.SavePNG(filename)
}
