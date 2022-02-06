package sk2d

import (
	"encoding/xml"
	"image"
	"log"
	"os"
)

func (p *Project) ExportXML(filename string) error {
	xmlFile, err := os.Create(filename)
	if err != nil {
		return err
	}

	enc := xml.NewEncoder(xmlFile)
	return enc.Encode(*p)
}

func LoadFile(filename string) *Project {
	p, err := parseXML(filename)
	if err != nil {
		log.Fatalf("failed to parse %s as project: %v", filename, err)
	}

	return p
}

func parseXML(filename string) (*Project, error) {
	p := New("", 0, 0)

	xmlFile, err := os.Open(filename)
	if err != nil {
		return p, err
	}

	doc := xml.NewDecoder(xmlFile)
	if err := doc.Decode(p); err != nil {
		return p, err
	}

	// boneCache:  make(map[string]int),
	// jointCache: make(map[string]string),
	// imageCache: make(map[string]image.Image),
	// animCache:  make(map[string]int),
	// skinCache:  make(map[string]int),

	// + Project.boneCache
	p.boneCache = make(map[string]int)
	for k, b := range p.Bones {
		p.boneCache[b.Name] = k
	}

	// + Project.jointCache
	// + Bone.jointCache
	p.jointCache = make(map[string]string)
	for _, b := range p.Bones {
		bone := p.getBone(b.Name)
		bone.jointCache = make(map[string]int)
		for k, joint := range b.Joints {
			p.jointCache[joint.Name] = b.Name
			bone.jointCache[joint.Name] = k
		}
	}

	// + Project.animCache
	p.animCache = make(map[string]int)
	for k, anim := range p.Anims {
		p.animCache[anim.Name] = k
	}

	// + Anim.frameCache
	for aName := range p.animCache {
		anim := p.getAnim(aName)
		anim.frameCache = make(map[string]int)
		for k, frame := range anim.Frames {
			anim.frameCache[frame.Name] = k
		}
	}

	// + Project.skinCache
	p.skinCache = make(map[string]int)
	for k, skin := range p.Skins {
		p.skinCache[skin.Name] = k
	}

	// + Project.imageCache
	p.imageCache = make(map[string]image.Image)
	for _, b := range p.Bones {
		if b.Image.Src != "" {
			p.mustLoadImage(b.Image.Src)
		}
	}
	for _, s := range p.Skins {
		for _, img := range s.Images {
			if img.Src != "" {
				p.mustLoadImage(img.Src)
			}
		}
	}

	return p, nil
}
