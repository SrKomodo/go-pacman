package main

import (
	"fmt"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const width = 144.0
const height = 168.0

const aspectRatio = 144.0 / 168.0

var lvl *level

func init() {
	runtime.LockOSThread()
}

func main() {
	lvl = newLevel("res/map.txt")

	// Init GLFW
	if err := glfw.Init(); err != nil {
		panic(fmt.Errorf("could not initialize glfw: %v", err))
	}
	defer glfw.Terminate()

	// Set window settings
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	// Create window
	win, err := glfw.CreateWindow(int(width*2), int(height*2), "Go Pacman", nil, nil)
	if err != nil {
		panic(fmt.Errorf("could not create glfw window: %v", err))
	}
	win.MakeContextCurrent()

	// Init OpenGL
	if err := gl.Init(); err != nil {
		panic(err)
	}
	gl.ClearColor(0, 0, 0, 1)

	// Correctly rezise OpenGL viewport
	win.SetFramebufferSizeCallback(func(window *glfw.Window, vW, vH int) {
		w := float64(vW)
		h := float64(vH)
		var newW float64
		var newH float64
		if w/h > aspectRatio {
			newW = h * aspectRatio
			newH = h
		} else {
			newW = w
			newH = w / aspectRatio
		}
		gl.Viewport(int32(w/2-newW/2), int32(h/2-newH/2), int32(newW), int32(newH))
	})

	bg := newSprite(
		"res/shaders/vert/default.glsl",
		"res/shaders/frag/bg.glsl",
		"res/textures/map.png",
		-1, -1, 2, 2,
		nil,
	)

	pacman := newPacman()
	prevT := 0.0

	// Main loop
	for !win.ShouldClose() {
		t := glfw.GetTime()
		dT := t - prevT

		glfw.PollEvents()

		gl.Clear(gl.COLOR_BUFFER_BIT)

		pacman.update(t, dT)
		bg.draw()
		pacman.draw()

		win.SwapBuffers()

		prevT = t
	}
}
