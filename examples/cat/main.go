package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bin16/sk2d"
)

func main() {
	genCat()
}

func genCat() {
	p := sk2d.LoadFile("cat.xml")
	p.SetSkin("main")
	p.SetAnim("main")
	p.SaveAs("cat.png")

	// you need imagemagick to go on

	sh := []string{
		"#!/bin/sh",
		"go run .",
	}

	for _, skin := range p.Skins {
		p.SetSkin(skin.Name)
		for _, anim := range p.Anims {
			p.SetAnim(anim.Name)
			for k, frame := range anim.Frames {
				p.SetFrame(frame.Name)

				p.Debug(false)
				filename1 := fmt.Sprintf("%s-%s-%s--%03d.tmp.png", p.Name, skin.Name, anim.Name, k)
				p.SaveAs(filename1)

				p.Debug(true)
				filename2 := fmt.Sprintf("%s-%s-%s--%03d.debug.png", p.Name, skin.Name, anim.Name, k)
				p.SaveAs(filename2)

				sh = append(sh, fmt.Sprintf("montage -geometry 256x256+8+8 %s %s %s", filename1, filename2, filename2))
				sh = append(sh, fmt.Sprintf("convert %s -bordercolor white -border 8 %s", filename2, filename2))

				sh = append(sh, fmt.Sprintf("montage -geometry 256x256+8+8 %s-*-%s--%03d.tmp.png montage_%03d.png", p.Name, anim.Name, k, k))
			}

			dstName := fmt.Sprintf("%s_%s_%s.gif", p.Name, skin.Name, anim.Name)
			debugDstName := fmt.Sprintf("%s_%s_%s_debug.gif", p.Name, skin.Name, anim.Name)

			sh = append(sh, fmt.Sprintf("convert %s-%s-%s--*.tmp.png %s", p.Name, skin.Name, anim.Name, dstName))
			sh = append(sh, fmt.Sprintf("convert %s-%s-%s--*.debug.png %s", p.Name, skin.Name, anim.Name, debugDstName))
		}
	}

	sh = append(sh, fmt.Sprintf("convert montage_*.png %s_montage.gif", p.Name))
	sh = append(sh, fmt.Sprintf("convert %s_montage.gif -bordercolor white -border 8 %s_montage.gif", p.Name, p.Name))

	sh = append(sh, "rm *.tmp.png")
	sh = append(sh, "rm *.debug.png")
	sh = append(sh, "rm montage_*.png")

	shFile, _ := os.Create("cat.sh")
	shFile.WriteString(strings.Join(sh, "\n"))
	shFile.Close()

}

func genBigDog() {
	p := sk2d.LoadFile("big-dog.xml")
	// p.Debug()

	for _, skin := range p.Skins {
		p.SetSkin(skin.Name)
		for _, anim := range p.Anims {
			p.SetAnim(anim.Name)
			for _, frame := range anim.Frames {
				p.SetFrame(frame.Name)
				filename := fmt.Sprintf("results/big-dog-%s-%s-%s.png", skin.Name, anim.Name, frame.Name)
				p.SaveAs(filename)
			}

			if anim.Name == "walk" && skin.Name == "main" {
				p.Debug(true)
				for _, frame := range anim.Frames {
					p.SetFrame(frame.Name)
					filename := fmt.Sprintf("results/debug-%s-%s-%s.png", skin.Name, anim.Name, frame.Name)
					p.SaveAs(filename)
				}
				p.Debug(false)
			}
		}
	}

	p.SaveAs("big-dog.png")
	p.SetSkin("yellow")
	p.SaveAs("big-dog-wood.png")

	p.SetAnim("walk")

	p.DebugAnim("walk")

	p.SetFrame("walk-1")
	p.SaveAs("big-dog-wood-frame-1.png")

	p.SetFrame("walk-2")
	p.SaveAs("big-dog-wood-frame-2.png")

	p.SetFrame("walk-3")
	p.SaveAs("big-dog-wood-frame-3.png")

	p.SetFrame("walk-4")
	p.SaveAs("big-dog-wood-frame-4.png")
}
