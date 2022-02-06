package sk2d

func (p *Project) mkBoneCache() {
	p.boneCache = make(map[string]int)
	for k, b := range p.Bones {
		p.boneCache[b.Name] = k
	}
}
