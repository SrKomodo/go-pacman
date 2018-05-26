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

	pacman := newSprite(
		"shaders/vertex.glsl",
		"shaders/pacman.glsl",
		-.5, .5, 1, 1,
		[]string{
			"t",
			"dir",
		},
	)

	var dir float32
	win.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		switch key {
		case glfw.KeyA:
			dir = 0
		case glfw.KeyW:
			dir = 1
		case glfw.KeyD:
			dir = 2
		case glfw.KeyS:
			dir = 3
		}
	})

	for !win.ShouldClose() {
		glfw.PollEvents()

		pacman.setUniform("t", float32(glfw.GetTime()))
		pacman.setUniform("dir", dir)

		gl.Clear(gl.COLOR_BUFFER_BIT)

		pacman.draw()

		win.SwapBuffers()
	}
}
