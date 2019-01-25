package main

import (
	"fmt"

	"github.com/thegtproject/gravity"
)

func handleInput(dt float32) {
	if checkDebugCommand() {
		return
	}

	if gravity.Pressed(gravity.KeyEscape) {
		gravity.Stop()
	}
	if gravity.Pressed(gravity.Key1) {
		terrain.Base().Transformer.RotateZ(dt * 7)
	}
	if gravity.Pressed(gravity.Key2) {
		terrain.Base().Transformer.RotateZ(dt * -7)
	}
	if gravity.Pressed(gravity.Key3) {
		//terrain.Transform(mgl32.Translate3D(0, 0, dt*20))
	}
	if gravity.Pressed(gravity.Key4) {
		//terrain.Transform(mgl32.Translate3D(0, 0, -dt*20))
	}

	if gravity.Pressed(gravity.KeyW) {
		cam.MoveForward(dt * 80)
	}
	if gravity.Pressed(gravity.KeyS) {
		cam.MoveBackward(dt * 80)

	}
	if gravity.Pressed(gravity.KeyA) {
		cam.MoveLeft(dt * 80)
	}
	if gravity.Pressed(gravity.KeyD) {
		cam.MoveRight(dt * 80)
	}
	if gravity.Pressed(gravity.KeySpace) {
		cam.MoveUp(dt * 80)
	}
	if gravity.Pressed(gravity.KeyLeftControl) {
		cam.MoveDown(dt * 80)
	}

	if gravity.Pressed(gravity.KeyLeft) {
		cam.Turn(dt * 30)
	}
	if gravity.Pressed(gravity.KeyRight) {
		cam.Turn(-dt * 30)
	}
	if gravity.Pressed(gravity.KeyUp) {
		cam.Roll(dt * 30)
	}
	if gravity.Pressed(gravity.KeyDown) {
		cam.Roll(-dt * 30)
	}

	// if gravity.Pressed(gravity.KeyLeft) {
	// 	terrain.Base().Transformer.RotateX(dt * 7)
	// }
	// if gravity.Pressed(gravity.KeyRight) {
	// 	terrain.Base().Transformer.RotateX(dt * -7)
	// }
	// if gravity.Pressed(gravity.KeyUp) {
	// 	terrain.Base().Transformer.RotateY(dt * 7)
	// }
	// if gravity.Pressed(gravity.KeyDown) {
	// 	terrain.Base().Transformer.RotateY(dt * -7)
	// }

	if gravity.JustPressed(gravity.KeyGrave) {
		debugCommandMode = true
		fmt.Print("debug command: ")
	}
}
