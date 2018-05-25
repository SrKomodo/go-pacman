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

	win, err := glfw.CreateWindow(600, 600, "Go Pacman", nil, nil)
	if err != nil {
		panic(fmt.Errorf("could not create glfw window: %v", err))
	}

	win.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	win.SetFramebufferSizeCallback(func(w *glfw.Window, width, height int) {
		if height < width {
			gl.Viewport(int32(width/2-height/2), 0, int32(height), int32(height))
		} else {
			gl.Viewport(0, int32(height/2-width/2), int32(width), int32(width))
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

	gl.ClearColor(0, 0.0, 1.0, 1.0)

	for !win.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)

		pacman.draw()

		win.SwapBuffers()
		glfw.PollEvents()
	}
}
