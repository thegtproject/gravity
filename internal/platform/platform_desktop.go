//+build !js

package platform

import (
	"runtime"

	"github.com/go-gl/gl/v4.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/thegtproject/gravity/internal/gravitygl"
	"github.com/thegtproject/gravity/internal/gravitygl/opengl"
	"github.com/thegtproject/gravity/internal/input"
)

type platformDesktop struct {
	win                 *glfw.Window
	mainThreadCallQueue chan func()
	cnt                 int
}

const maxQueuedCalls = 8

func (dt *platformDesktop) getMainThreadCallQueue() chan func() {
	println("Platform<Desktop>.getMainThreadCallQueue()")
	return dt.mainThreadCallQueue
}

func newPlatform(title string, width int, height int) *platformDesktop {
	println("Platform<Desktop>.newPlatform()")
	if err := glfw.Init(); err != nil {
		panic("failed to initialize glfw: " + err.Error())
	}
	initGLFW()
	win := createWindow(title, width, height)

	platform := &platformDesktop{win: win, mainThreadCallQueue: make(chan func(), maxQueuedCalls)}
	platform._keyCallbackHandler()
	platform._mousePositionCallbackHandler()
	platform._mouseButtonCallbackHandler()
	return platform
}

func createWindow(title string, width int, height int) *glfw.Window {
	println("Platform<Desktop>.createWindow()")
	win, err := glfw.CreateWindow(
		width,
		height,
		title,
		nil,
		nil,
	)
	if err != nil {
		panic(err)
	}
	win.MakeContextCurrent()
	return win
}

func initGLFW() {
	println("Platform<Desktop>.initGLFW()")
	if err := glfw.Init(); err != nil {
		panic("failed to initialize glfw: " + err.Error())
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
}

// SetKeyCallbackFunction ...
func (dt *platformDesktop) SetKeyCallbackFunction(f func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey)) {
	println("Platform<Desktop>.SetKeyCallbackFunction()")
	dt.win.SetKeyCallback(f)
}

// OnResize ...
func (dt *platformDesktop) OnResize(f func(width, height int)) {
	println("Platform<Desktop>.OnResize()")
	dt.win.SetSizeCallback(
		func(w *glfw.Window, width, height int) {
			f(width, height)
		},
	)
}

// Update ...
func (dt *platformDesktop) Update() {
	dt.mainThreadCallQueue <- func() {
		dt.win.SwapBuffers()
		glfw.PollEvents()
		gl.Clear(gl.COLOR_BUFFER_BIT)
	}
}

// SetClearColor ...
func (dt *platformDesktop) SetClearColor(col mgl32.Vec4) {
	gravitygl.SetClearColor(col)
}

// GetUnderlyingPlatformObject ...
func (dt *platformDesktop) GetUnderlyingPlatformObject() *glfw.Window {
	println("Platform<Desktop>.GetUnderlyingPlatformObject()")
	return dt.win
}

// Running ...
func (dt *platformDesktop) Running() bool {
	return !dt.win.ShouldClose()
}

func (dt *platformDesktop) defaults() {
	// TODO: Find a better spot for this. This is temporary.
	gravitygl.Enable(opengl.DEPTH_TEST)
	gravitygl.DepthFunc(opengl.LEQUAL)
	gravitygl.Enable(opengl.BLEND)
	gravitygl.BlendFunc(opengl.ONE, opengl.ONE_MINUS_SRC_ALPHA)
	println("Platform<Desktop>.defaults() executed")
}

func (dt *platformDesktop) Run(run func()) {
	println("Platform<Desktop>.Run()")

	println("Platform<Desktop>.defaults() queued")
	dt.defaults()

	done := make(chan struct{}, 1)
	go func() {
		run()
		done <- struct{}{}
	}()

	running := true
	for running == true {
		select {
		case mtcall := <-dt.mainThreadCallQueue:
			mtcall()
		case <-done:
			running = false
		}
	}

	println("Platform<Desktop>.Run() Finished!")
}

func (dt *platformDesktop) Stop() {
	println("Platform<Desktop>.Stop()")
	dt.win.SetShouldClose(true)
}

func (dt *platformDesktop) _keyCallbackHandler() {
	println("Platform<Desktop>._keyCallbackHandler() queued")
	go func() {
		dt.mainThreadCallQueue <- func() {
			println("Platform<Desktop>._keyCallbackHandler() executed")
			dt.win.SetKeyCallback(input.KeyCallbackHandler)
		}
	}()
}

func (dt *platformDesktop) _mouseButtonCallbackHandler() {
	println("Platform<Desktop>._mouseButtonCallbackHandler() queued")
	go func() {
		dt.mainThreadCallQueue <- func() {
			println("Platform<Desktop>._mouseButtonCallbackHandler() executed")
			dt.win.SetMouseButtonCallback(input.MouseButtonCallbackHandler)
		}
	}()
}

func (dt *platformDesktop) _mousePositionCallbackHandler() {
	println("Platform<Desktop>._mousePositionCallbackHandler() queued")
	go func() {
		dt.mainThreadCallQueue <- func() {
			println("Platform<Desktop>._mousePositionCallbackHandler() executed")
			dt.win.SetCursorPosCallback(input.MousePositionHandler)
		}
	}()
}

func init() {
	runtime.LockOSThread()
}
