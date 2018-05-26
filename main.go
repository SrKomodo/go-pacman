package main

import (
	"fmt"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

var width = 144.0 * 2
var height = 168.0 * 2

const aspectRatio = 144.0 / 168.0

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

	win, err := glfw.CreateWindow(int(width), int(height), "Go Pacman", nil, nil)
	if err != nil {
		panic(fmt.Errorf("could not create glfw window: %v", err))
	}

	win.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}
	gl.ClearColor(0, 0, 0, 1)

	win.SetFramebufferSizeCallback(func(win *glfw.Window, newW, newH int) {
		w := float64(newW)
		h := float64(newH)
		if w/h > aspectRatio {
			width = h * aspectRatio
			height = h
		} else {
			width = w
			height = w / aspectRatio
		}
		gl.Viewport(
			int32(w/2-width/2),
			int32(h/2-height/2),
			int32(width),
			int32(height),
		)
	})

	bg := newSprite(
		"shaders/vert/default.glsl",
		"shaders/frag/bg.glsl",
		"textures/map.png",
		-1, 1, 2, 2,
		nil,
	)

	pacman := newPacman()

	win.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		pacman.onKey(key)
	})

	for !win.ShouldClose() {
		glfw.PollEvents()

		pacman.update()

		gl.Clear(gl.COLOR_BUFFER_BIT)

		bg.draw()
		pacman.draw()

		win.SwapBuffers()
	}
}
