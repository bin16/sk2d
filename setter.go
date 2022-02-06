package sk2d

import (
	"log"
	"strconv"
)

type Setter struct {
	Subject string `xml:"subject,attr"`
	Target  string `xml:"target,attr"`
	Name    string `xml:"name,attr"`
	Content string `xml:"content,attr"`
}

func (s *Setter) ParseFloat64() float64 {
	value, err := strconv.ParseFloat(s.Content, 64)
	if err != nil {
		log.Panicf("<frame.setter> failed to parse [%s] as float64\n", s.Content)
	}

	return value
}

func (s *Setter) ParseInt() int {
	value, err := strconv.ParseInt(s.Content, 10, 64)
	if err != nil {
		log.Panicf("<frame.setter> failed to parse [%s] as int\n", s.Content)
	}

	return int(value)
}

func (p *Project) getBoneSetter(bName, name string) *Setter {
	f := p.getFrame(p.anim, p.frame)
	if f == nil {
		return nil
	}

	for _, s := range f.Setters {
		if s.Subject == "bone" && s.Target == bName && s.Name == name {
			return &s
		}
	}

	return nil
}

func (p *Project) getJointSetter(jName, name string) *Setter {
	skin := p.getSkin(p.skin)
	if skin == nil {
		return nil
	}

	for _, s := range skin.Setters {
		if s.Subject == "bone.joint" && s.Target == jName && s.Name == name {
			return &s
		}
	}

	return nil
}

func (p *Project) getRootSetter(name string) *Setter {
	f := p.getFrame(p.anim, p.frame)
	if f == nil {
		return nil
	}

	for _, s := range f.Setters {
		if s.Subject == "root" && s.Target == "root" && s.Name == name {
			return &s
		}
	}

	return nil
}
