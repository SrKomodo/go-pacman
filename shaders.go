package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

func compileShader(path string, shaderType uint32) uint32 {
	// Read shader file
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	data := string(bytes) + "\x00" // Dont forget the null terminators!

	shader := gl.CreateShader(shaderType)   // Create shader
	source, free := gl.Strs(data)           // Do OpenGL dark magic
	gl.ShaderSource(shader, 1, source, nil) // Load the shader
	free()                                  // More OpenGL dark magic
	gl.CompileShader(shader)                // Compile the shader

	// Check if shader compiled correctly
	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var size int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &size)

		log := strings.Repeat("\x00", int(size+1))
		gl.GetShaderInfoLog(shader, size, nil, gl.Str(log))

		panic(fmt.Errorf("failed to compile shader: %v", log))
	}

	return shader
}

func createProgram(vertexPath, fragmentPath string) uint32 {
	// Compile shaders
	vertex := compileShader(vertexPath, gl.VERTEX_SHADER)
	fragment := compileShader(fragmentPath, gl.FRAGMENT_SHADER)

	program := gl.CreateProgram() // Create program
	// Attach shaders
	gl.AttachShader(program, vertex)
	gl.AttachShader(program, fragment)

	gl.BindFragDataLocation(program, 0, gl.Str("gl_Color\x00")) // Link `c` attribute in fragment shader

	gl.LinkProgram(program) // Compile program

	// Check if program compiled correctly
	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		panic(fmt.Errorf("failed to link program: %v", log))
	}

	return program
}
