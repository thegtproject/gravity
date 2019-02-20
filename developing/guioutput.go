package main

import (
	"bytes"
	"fmt"

	imgui "github.com/thegtproject/gravity/internal/imgui-go"
)

type guiOutput struct {
	buf           bytes.Buffer
	posOpened     imgui.Vec2
	posCollapsed  imgui.Vec2
	sizeOpened    imgui.Vec2
	sizeCollapsed imgui.Vec2
	opened        bool
}

// GUIOutput ...
var GUIOutput = guiOutput{
	posOpened:     imgui.Vec2{X: 505, Y: 768 - 235},
	sizeOpened:    imgui.Vec2{X: 856, Y: 200},
	posCollapsed:  imgui.Vec2{X: 505, Y: 768 - 235 + 185},
	sizeCollapsed: imgui.Vec2{X: 856, Y: 15},
}

func (g *guiOutput) Print(a ...interface{}) {
	fmt.Fprint(&g.buf, a...)
}

func (g *guiOutput) Printf(format string, a ...interface{}) {
	fmt.Fprintf(&g.buf, format, a...)
}

func (g *guiOutput) Println(a ...interface{}) {
	fmt.Fprintln(&g.buf, a...)
}

func (g *guiOutput) Render() {
	const flags = 0 |
		imgui.WindowFlagsNoResize |
		imgui.WindowFlagsNoSavedSettings

	if g.opened {
		imgui.SetNextWindowPos(g.posOpened)
		imgui.SetNextWindowSize(g.sizeOpened)
	} else {

		imgui.SetNextWindowPos(g.posCollapsed)
		imgui.SetNextWindowSize(g.sizeCollapsed)
	}
	imgui.PushStyleVarFloat(imgui.StyleVarWindowRounding, 0)
	imgui.PushStyleColor(imgui.StyleColorButton, colorCollapseBtn)
	imgui.PushStyleColor(imgui.StyleColorButtonActive, colorCollapseBtn)
	imgui.PushStyleColor(imgui.StyleColorButtonHovered, colorCollapseBtnHovered)

	g.opened = imgui.BeginV("output", nil, flags)

	imgui.Text(g.buf.String())
	imgui.SetScrollHereY(1.0)
	imgui.End()
	imgui.PopStyleVar()
	imgui.PopStyleColorV(3)
}
