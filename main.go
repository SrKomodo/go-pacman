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

	// Create VAO
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	// Create VBO
	vertices := []float32{
		0, 0.5,
		.5, -.5,
		-.5, -.5,
	}

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4 /* I have no idea why i have to multiply */, gl.Ptr(vertices), gl.STATIC_DRAW)

	// Create program
	program := createProgram("shaders/vertex.glsl", "shaders/fragment.glsl")

	gl.UseProgram(program)
	gl.ClearColor(0, 0.0, 1.0, 1.0)

	for !win.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.DrawArrays(gl.TRIANGLES, 0, 3)

		win.SwapBuffers()
		glfw.PollEvents()
	}
}
