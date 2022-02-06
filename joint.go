package sk2d

import (
	"fmt"
)

type HAlignValue string
type VAlignValue string
type JointPosition string

const (
	HAlignStart  HAlignValue = "start"
	HAlignCenter HAlignValue = "center"
	HAlignEnd    HAlignValue = "end"

	VAlignTop    VAlignValue = "top"
	VAlignMiddle VAlignValue = "middle"
	VAlignBottom VAlignValue = "bottom"

	Absolute JointPosition = "absolute"
	Relative JointPosition = "relative"
)

type Joint struct {
	Name    string      `xml:"name,attr"`
	OffsetX int         `xml:"x,attr"`
	OffsetY int         `xml:"y,attr"`
	VAlign  VAlignValue `xml:"valign,attr"`
	HAlign  HAlignValue `xml:"halign,attr"`
}

func (p *Project) AddJoint(boneName, jointName string) {
	b := p.getBone(boneName)
	b.Joints = append(b.Joints, Joint{
		Name: jointName,
	})

	b.jointCache = make(map[string]int)
	for k, j := range b.Joints {
		b.jointCache[j.Name] = k
	}

	p.jointCache[jointName] = boneName
}

func (p *Project) SetJointOffsetX(jName string, x int, halign HAlignValue) {
	j := p.getJoint(jName)
	j.HAlign = halign
	j.OffsetX = x
}

func (p *Project) SetJointOffsetY(jName string, y int, valign VAlignValue) {
	j := p.getJoint(jName)
	j.VAlign = valign
	j.OffsetY = y
}

func (p *Project) Join(jFrom, jTo string, position JointPosition) {
	bName := p.jointCache[jFrom]
	b := p.getBone(bName)
	b.Map.From = jFrom
	b.Map.To = jTo
	b.Map.Position = position

	if p.debug {
		fmt.Printf("Join %s, %s, %v\n", jFrom, jTo, b.Map)
	}
}

// TODO: check joint exists
func (p *Project) getJoint(jName string) *Joint {
	bName, ok := p.jointCache[jName]
	if !ok {
		return nil
	}

	b := p.getBone(bName)
	if b == nil {
		return nil
	}

	d, ok := b.jointCache[jName]
	if !ok {
		return nil
	}

	return &b.Joints[d]
}
