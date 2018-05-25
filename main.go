package main

import (
	"fmt"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const width = 600
const height = 600

func init() {
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		panic(fmt.Errorf("could not initialize glfw: %v", err))
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	win, err := glfw.CreateWindow(width, height, "Go Pacman", nil, nil)
	if err != nil {
		panic(fmt.Errorf("could not create glfw window: %v", err))
	}

	win.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}
	gl.ClearColor(0, 0.0, 1.0, 1.0)

	win.SetFramebufferSizeCallback(func(win *glfw.Window, newW, newH int) {
		if newH < newW {
			gl.Viewport(int32(newW/2-newH/2), 0, int32(newH), int32(newH))
		} else {
			gl.Viewport(0, int32(newH/2-newW/2), int32(newW), int32(newW))
		}
	})

	pacman := newEntity(
		"shaders/vertex.glsl",
		"shaders/pacman.glsl",
		createSquare(-.5, -.5, 1, 1),
		[]attribute{
			{"p\x00", 2},
			{"_uv\x00", 2},
		},
	)

	t := gl.GetUniformLocation(pacman.program, gl.Str("t\x00"))

	for !win.ShouldClose() {
		gl.Uniform1f(t, float32(glfw.GetTime()))

		gl.Clear(gl.COLOR_BUFFER_BIT)

		pacman.draw()

		win.SwapBuffers()
		glfw.PollEvents()
	}
}
