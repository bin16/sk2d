package sk2d

import (
	"fmt"
	"log"
)

func (p *Project) DebugAnim(name string) {
	anim := p.getAnim(name)
	fmt.Printf("[ DebugAnim ]\n")
	fmt.Printf("[ Anim ] %s\n", name)
	fmt.Printf("[ Anim.Frames ]\n")
	fmt.Printf("[ getAnim() ] %v\n", p.getAnim(name))
	for _, frame := range anim.Frames {
		p.SetFrame(frame.Name)
		fmt.Printf("[ Frame ] %s\n", frame.Name)

		fmt.Printf("[ getFrame() ] %v\n", p.getFrame(name, frame.Name))

		for _, setter := range frame.Setters {
			fmt.Printf("[ Setter ] %s\n", setter.Name)
			if setter.Name == "angle" {
				b := p.getBone(setter.Target)
				fmt.Printf("[ Bone.Angle ] %f\n", b.Angle)
				fmt.Printf("[ Setter.Content ] %s\n", setter.Content)
				fmt.Printf("[ getBoneAngle() ] %f\n", p.getBoneAngle(b.Name))
				fmt.Printf("[ getBoneSetter() ] %v\n", p.getBoneSetter(b.Name, "angle"))
				fmt.Printf("[ anim ] %s\n", p.anim)
				fmt.Printf("[ frame ] %s\n", p.frame)
				fmt.Printf("[ getFrame() ] %v\n", p.getFrame(p.anim, p.frame))
			}
		}
	}
}

func (p *Project) DebugBone(bName string) {
	b := p.getBone(bName)
	fmt.Printf("\n==== DebugBone, [%s] ====\n", bName)
	fmt.Printf("[bone.map] [%s] => [%s], %s\n", b.Map.From, b.Map.To, b.Map.Position)

	d := p.boneCache[bName]
	fmt.Printf("[boneCache]: [%s] => [%d]\n", bName, d)
	for jName, name := range p.jointCache {
		if name == bName {
			fmt.Printf("* jointCache: [%s] => [%s]\n", jName, bName)
		}
	}

	fmt.Println("[bone.joints]")
	for _, j := range b.Joints {
		fmt.Printf("* joint: [%s] (%d, %d) ", j.Name, j.OffsetX, j.OffsetY)
		x, y := p.getJointLocation(j.Name)
		fmt.Printf("{%.2f, %.2f}\n", x, y)
	}

	fmt.Println("[bone.angle]")
	fmt.Printf("angle=%.0f°\n", p.getBoneAngle(bName))

	fmt.Printf("==== DebugBone, [%s] ====\n\n", bName)
}

func (p *Project) DebugJoint(jName string) {
	x, y := p.getJointLocation(jName)
	log.Println(x, y, jName, "[ DebugJoint ]")
}

func (p *Project) Debug(b bool) {
	p.debug = b
}

func (p *Project) Check() {
	p.getBone("head-1")

	p.DebugBones()

	p.getJoint("head-1-j")
}

func (p *Project) DebugBones() {
	for _, b := range p.Bones {
		bw, bh := p.getBoneSize(b.Name)
		fmt.Printf("[ bone ] %s (%dx%d)\n", b.Name, bw, bh)
		for _, j := range b.Joints {
			x, y := p.boneGetJointLocation(b.Name, j.Name)
			fmt.Printf("[ bone.joint ] %s [halign=%6s, valign=%6s] (%d, %d)\n", j.Name, j.HAlign, j.VAlign, x, y)
		}

		for jName, d := range b.jointCache {
			j := b.Joints[d]
			fmt.Printf("[ Bone.jointCache ] <%s> => <%d> => <%s>\n", jName, d, j.Name)
		}

		for _, anim := range p.Anims {
			fmt.Printf("[ animation ] %s\n", anim.Name)
			for _, frame := range anim.Frames {
				fmt.Printf("[ animation.frame ] %s\n", frame.Name)
				for _, setter := range frame.Setters {
					fmt.Printf("[ setter ] %s, %s, %s, %s\n", setter.Subject, setter.Target, setter.Name, setter.Content)
				}
			}
		}
	}

	for jName, bName := range p.jointCache {
		fmt.Printf("[ Project.jointCache ] <%s> => <%s>\n", jName, bName)
	}

	// for jName, _ := range p.jointCache {
	// 	j := p.getJoint(jName)
	// 	fmt.Printf("[ getJoint ] <%s> => <%s>\n", jName, j.Name)
	// }
	// for bName := range p.boneCache {
	// 	b := p.getBone(bName)
	// 	fmt.Printf("[ getBone ] <%s> => <%s>\n", bName, b.Name)
	// }

	fmt.Println()
}
