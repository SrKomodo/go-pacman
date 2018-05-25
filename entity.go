package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type attribute struct {
	name string
	size int32
}

type entity struct {
	vao      uint32
	program  uint32
	vertices []float32
}

func (e *entity) draw() {
	gl.BindVertexArray(e.vao)
	gl.UseProgram(e.program)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(e.vertices)/2))
}

func newEntity(vertexPath, fragmentPath string, vertices []float32, attributes []attribute) *entity {
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
	var totalSize int32
	for _, attrib := range attributes {
		totalSize += attrib.size
	}

	offset := 0
	for _, attrib := range attributes {
		pointer := uint32(gl.GetAttribLocation(program, gl.Str(attrib.name)))
		gl.EnableVertexAttribArray(pointer)
		gl.VertexAttribPointer(pointer, attrib.size, gl.FLOAT, false, totalSize-attrib.size, gl.PtrOffset(offset))
	}

	return &entity{vao, program, vertices}
}
