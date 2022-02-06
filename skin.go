package sk2d

type Skin struct {
	Name    string      `xml:"name,attr"`
	Images  []SkinImage `xml:"skin.images>image"`
	Setters []Setter    `xml:"skin.setters>setter"`
}

func (p *Project) AddSkin(name string) {
	if s := p.getSkin(name); s != nil {
		return
	}

	p.Skins = append(p.Skins, Skin{
		Name:   name,
		Images: []SkinImage{},
	})
	for k, skin := range p.Skins {
		p.skinCache[skin.Name] = k
	}
}

func (p *Project) SetSkin(skin string) {
	p.skin = skin
}

func (p *Project) getSkin(name string) *Skin {
	if _, ok := p.skinCache[name]; !ok {
		return nil
	}

	d := p.skinCache[name]
	return &p.Skins[d]
}

func (p *Project) SetSkinImage(skinName, bName, imgSrc string) {
	s := p.getSkin(skinName)
	if s == nil {
		p.AddSkin(skinName)
	}
	s = p.getSkin(skinName)

	s.Images = append(s.Images, SkinImage{
		Target: bName,
		Src:    imgSrc,
	})
	p.mustLoadImage(imgSrc)
}
