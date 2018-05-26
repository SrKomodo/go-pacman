package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

func newShader(path string, shaderType uint32) uint32 {
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

func newProgram(vertexPath, fragmentPath string) uint32 {
	// Compile shaders
	vertex := newShader(vertexPath, gl.VERTEX_SHADER)
	fragment := newShader(fragmentPath, gl.FRAGMENT_SHADER)

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

func newTexture(path string) uint32 {
	imgFile, err := os.Open(path)
	if err != nil {
		panic(fmt.Errorf("could not open %s: %v", path, err))
	}

	img, err := png.Decode(imgFile)
	if err != nil {
		panic(fmt.Errorf("could not decode %s: %v", path, err))
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		panic("unsuported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var texture uint32
	gl.GenTextures(1, &texture)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix),
	)

	return texture
}
