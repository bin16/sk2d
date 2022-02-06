package sk2d

import (
	"log"
	"math"
)

type Bone struct {
	Name   string `xml:"name,attr"`
	Width  int    `xml:"width,attr"`
	Height int    `xml:"height,attr"`

	Angle  float64 `xml:"angle,attr"`
	ZIndex int     `xml:"z-index,attr"`

	Joints []Joint   `xml:"bone.joints>joint"`
	Map    BoneMap   `xml:"bone.map"`
	Image  BoneImage `xml:"bone.image"`

	jointCache map[string]int
}

func (p *Project) AddBone(boneName string) {
	if b := p.getBone(boneName); b != nil {
		log.Panicf("bone [%s] exists", boneName)
	}

	p.Bones = append(p.Bones, Bone{
		Name:   boneName,
		Joints: []Joint{},
		Map:    BoneMap{},
	})
	p.mkBoneCache()
}

func (p *Project) Rotate(bName string, angle float64) {
	b := p.getBone(bName)
	b.Angle = angle
}

func (p *Project) getImageSize(src string) (w, h int) {
	img, ok := p.imageCache[src]
	if !ok {
		return 0, 0
	}

	w = img.Bounds().Dx()
	h = img.Bounds().Dy()
	return
}

func (p *Project) getBoneImageSrc(bName string) string {
	b := p.getBone(bName)
	s := p.getSkin(p.skin)
	if s != nil {
		for _, img := range s.Images {
			if img.Target == bName {
				return img.Src
			}
		}
	}

	return b.Image.Src
}

func (p *Project) getBoneSize(bName string) (w, h int) {
	b := p.getBone(bName)
	w = b.Width
	h = b.Height
	imgSrc := p.getBoneImageSrc(b.Name)
	imgW, imgH := p.getImageSize(imgSrc)
	if w == 0 {
		w = imgW
	}

	if h == 0 {
		h = imgH
	}

	return
}

func intsToFloat64(x, y int) (x1, y1 float64) {
	x1 = float64(x)
	y1 = float64(y)
	return
}

// TODO: check override
func (p *Project) getBoneAngle(bName string) float64 {
	b := p.getBone(bName)

	angle := b.Angle
	if s := p.getBoneSetter(bName, "angle"); s != nil {
		angle = s.ParseFloat64()
	}

	if b.Map.To == "root" {
		return angle
	}

	if b.Map.Position == Absolute {
		return angle
	}

	pBoneName := p.jointCache[b.Map.To]
	return angle + p.getBoneAngle(pBoneName)
}

func (p *Project) getAngle(bName string) float64 {
	ang := p.getBoneAngle(bName)
	return ang * math.Pi / 180
}

// TODO: check bone exists
func (p *Project) getBone(bName string) *Bone {
	d, ok := p.boneCache[bName]
	if !ok {
		return nil
	}

	if d > len(p.Bones) {
		return nil
	}

	return &p.Bones[d]
}

func (p *Project) boneGetPointRelLocation(bName string, x0, y0 int) (x, y int) {
	b := p.getBone(bName)

	jFromX, jFromY := p.boneGetJointLocation(bName, b.Map.From)
	x = x0 - jFromX
	y = y0 - jFromY
	return
}

// joint location (x, y int) to (bone.Map.From)
func (p *Project) boneGetJointRelLocation(bName, jName string) (x, y int) {
	b := p.getBone(bName)
	x0, y0 := p.boneGetJointLocation(bName, b.Map.From)
	x1, y1 := p.boneGetJointLocation(bName, jName)
	x = x1 - x0
	y = y1 - y0
	return
}

// joint location (x, y int) to bone (0, 0)
func (p *Project) boneGetJointLocation(bName, jName string) (x, y int) {
	j := p.getJoint(jName)
	bw, bh := p.getBoneSize(bName)

	// dx, dy are OffsetX, OffsetY with setters
	dx := j.OffsetX
	dy := j.OffsetY
	if s := p.getJointSetter(jName, "x"); s != nil {
		dx = s.ParseInt()
	}
	if s := p.getJointSetter(jName, "y"); s != nil {
		dy = s.ParseInt()
	}

	x = dx
	y = dx

	if j.HAlign == HAlignCenter {
		x = bw/2 + dx
	} else if j.HAlign == HAlignEnd {
		x = bw - dx
	} else {
		x = dx
	}

	if j.VAlign == VAlignMiddle {
		y = bh/2 + dy
	} else if j.VAlign == VAlignBottom {
		y = bh - dy
	} else {
		y = dy
	}

	// log.Println(bName, jName, j.HAlign, j.OffsetX, j.VAlign, j.OffsetY, x, y)

	return
}
