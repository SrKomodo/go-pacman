package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type sprite struct {
	vao      uint32
	program  uint32
	size     int32
	uniforms map[string]int32
}

func (e *sprite) draw() {
	gl.BindVertexArray(e.vao)
	gl.UseProgram(e.program)
	gl.DrawArrays(gl.TRIANGLE_STRIP, 0, e.size)
}

func (e *sprite) setUniform(name string, value float32) {
	gl.Uniform1f(e.uniforms[name], value)
}

func newSprite(
	vertexPath, fragmentPath string,
	x, y, w, h float32,
	uniforms []string,
) *sprite {
	// Generate vertices
	vertices := []float32{
		x, y - h, 0, 1,
		x, y, 0, 0,
		x + w, y - h, 1, 1,
		x + w, y, 1, 0,
	}

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	program := createProgram(vertexPath, fragmentPath)
	gl.UseProgram(program)

	// Link vertex attributes
	coords := uint32(gl.GetAttribLocation(program, gl.Str("coords\x00")))
	gl.EnableVertexAttribArray(coords)
	gl.VertexAttribPointer(coords, 2, gl.FLOAT, false, 4*4, gl.PtrOffset(0))

	uv := uint32(gl.GetAttribLocation(program, gl.Str("uv\x00")))
	gl.EnableVertexAttribArray(uv)
	gl.VertexAttribPointer(uv, 2, gl.FLOAT, false, 4*4, gl.PtrOffset(2*4))

	// Create uniform pointers
	newUniforms := make(map[string]int32)
	for _, uniform := range uniforms {
		newUniforms[uniform] = gl.GetUniformLocation(program, gl.Str(uniform+"\x00"))
	}

	return &sprite{vao, program, int32(len(vertices)) / 4, newUniforms}
}