package sk2d

import "fmt"

type BoneMap struct {
	From     string        `xml:"from,attr"`
	To       string        `xml:"to,attr"`
	Position JointPosition `xml:"position,attr"`
}

func (m *BoneMap) String() string {
	return fmt.Sprintf("[%s=>%s]", m.From, m.To)
}
