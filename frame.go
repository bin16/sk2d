package sk2d

import "fmt"

type Frame struct {
	Name    string   `xml:"name,attr"`
	Setters []Setter `xml:"frame.setters>setter"`
}

func (p *Project) AddFrame(aName string, fName string) {
	anim := p.getAnim(aName)
	anim.Frames = append(anim.Frames, Frame{
		Name:    fName,
		Setters: []Setter{},
	})
	for k, frame := range anim.Frames {
		anim.frameCache[frame.Name] = k
	}
}

func (p *Project) getFrame(aName, fName string) *Frame {
	anim := p.getAnim(aName)
	if anim == nil {
		return nil
	}

	d, ok := anim.frameCache[fName]
	if !ok {
		return nil
	}

	return &anim.Frames[d]
}

func (p *Project) SetAnim(anim string) {
	p.anim = anim
}

func (p *Project) SetFrame(frame string) {
	p.frame = frame
}

func (p *Project) SetBoneAngle(bName string, angle float64) {
	f := p.getFrame(p.anim, p.frame)
	f.Setters = append(f.Setters, Setter{
		Subject: "bone",
		Target:  bName,
		Name:    "angle",
		Content: fmt.Sprintf("%.6f", angle),
	})
}
