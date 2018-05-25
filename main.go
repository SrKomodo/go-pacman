package main

import (
	"fmt"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func init() {
	runtime.LockOSThread()
}

func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(fmt.Errorf("could not initialize glfw: %v", err))
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	win, err := glfw.CreateWindow(800, 600, "Go Pacman", nil, nil)
	if err != nil {
		panic(fmt.Errorf("could not create glfw window: %v", err))
	}

	win.MakeContextCurrent()

	return win
}

func main() {
	win := initGlfw()
	defer glfw.Terminate()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	square := newEntity(
		"shaders/vertex.glsl",
		"shaders/fragment.glsl",
		createSquare(-.5, -.5, .75, .75),
		[]attribute{
			{"p\x00", 2},
			{"_uv\x00", 2},
		},
	)

	otherSquare := newEntity(
		"shaders/vertex.glsl",
		"shaders/fragment2.glsl",
		createSquare(-.25, -.25, .75, .75),
		[]attribute{
			{"p\x00", 2},
			{"_uv\x00", 2},
		},
	)

	gl.ClearColor(0, 0.0, 1.0, 1.0)

	for !win.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)

		square.draw()
		otherSquare.draw()

		win.SwapBuffers()
		glfw.PollEvents()
	}
}
