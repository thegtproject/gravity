package main

import (
	"bytes"
	"fmt"

	imgui "github.com/thegtproject/gravity/internal/imgui-go"
)

type guiCamera struct {
	buf           bytes.Buffer
	posOpened     imgui.Vec2
	posCollapsed  imgui.Vec2
	sizeOpened    imgui.Vec2
	sizeCollapsed imgui.Vec2
	opened        bool
}

// GUICamera ...
var GUICamera = guiCamera{
	posOpened:  imgui.Vec2{X: 5, Y: 150},
	sizeOpened: imgui.Vec2{X: 250, Y: 225},
}

func (g *guiCamera) Print(a ...interface{}) {
	fmt.Fprint(&g.buf, a...)
}

func (g *guiCamera) Printf(format string, a ...interface{}) {
	fmt.Fprintf(&g.buf, format, a...)
}

func (g *guiCamera) Println(a ...interface{}) {
	fmt.Fprintln(&g.buf, a...)
}

func (g *guiCamera) Render() {
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

	g.opened = imgui.BeginV("camera", nil, flags)

	imgui.Text(g.buf.String())
	imgui.SetScrollHereY(1.0)
	imgui.End()
	imgui.PopStyleVar()
	imgui.PopStyleColorV(3)
}
