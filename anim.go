package sk2d

type Animation struct {
	Name   string  `xml:"name,attr"`
	Frames []Frame `xml:"animation.frames>frame"`

	frameCache map[string]int
}

func (p *Project) AddAnim(aName string) {
	if _, ok := p.animCache[aName]; ok {
		return
	}

	p.Anims = append(p.Anims, Animation{
		Name:       aName,
		Frames:     []Frame{},
		frameCache: make(map[string]int),
	})
	for k, anim := range p.Anims {
		p.animCache[anim.Name] = k
	}
}

func (p *Project) getAnim(aName string) *Animation {
	if _, ok := p.animCache[aName]; !ok {
		return nil
	}

	d, ok := p.animCache[aName]
	if !ok {
		return nil
	}

	return &p.Anims[d]
}
