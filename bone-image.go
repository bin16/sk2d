package sk2d

import (
	"image/png"
	"log"
	"os"
)

type BoneImage struct {
	Src string `xml:"src,attr"`
}

func (p *Project) SetBoneImage(bName, src string) {
	b := p.getBone(bName)
	b.Image.Src = src
	p.mustLoadImage(src)
}

// ================

func (p *Project) loadImage(src string) error {
	if _, ok := p.imageCache[src]; ok {
		return nil
	}

	imgFile, err := os.Open(src)
	if err != nil {
		return err
	}

	img, err := png.Decode(imgFile)
	if err != nil {
		return err
	}

	p.imageCache[src] = img
	return nil
}

func (p *Project) mustLoadImage(src string) {
	if err := p.loadImage(src); err != nil {
		log.Panicf("failed to load image, %v", err)
	}
}
