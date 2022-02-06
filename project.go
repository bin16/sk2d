package sk2d

import (
	"image"
)

type Project struct {
	Name    string `xml:"name,attr"`
	Width   int    `xml:"width,attr"`
	Height  int    `xml:"height,attr"`
	OffsetX int    `xml:"x,attr"`
	OffsetY int    `xml:"y,attr"`

	Fill   string `xml:"fill,attr"`
	Stroke string `xml:"stroke,attr"`

	Bones []Bone      `xml:"project.bones>bone"`
	Anims []Animation `xml:"project.animations>animation"`
	Skins []Skin      `xml:"project.skins>skin"`

	boneCache  map[string]int         // boneName => boneIndex
	animCache  map[string]int         // animName => animIndex
	jointCache map[string]string      // jointName => boneName
	skinCache  map[string]int         // skinName => skinIndex
	imageCache map[string]image.Image // imageSrc => image.Image

	debug bool

	anim  string
	frame string
	skin  string
}

func New(pName string, w, h int) *Project {
	p := &Project{
		Name:       pName,
		Width:      w,
		Height:     h,
		Bones:      []Bone{},
		Anims:      []Animation{},
		boneCache:  make(map[string]int),
		jointCache: make(map[string]string),
		imageCache: make(map[string]image.Image),
		animCache:  make(map[string]int),
		skinCache:  make(map[string]int),
	}

	return p
}
