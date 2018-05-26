package main

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

type pacman struct {
	sprite *sprite
	dir    float32
	x      float32
	y      float32
}

func (p *pacman) update() {
	p.sprite.setUniform("t", float32(glfw.GetTime()))
	p.sprite.setUniform("dir", p.dir)
	p.sprite.setUniform("x", p.x)
	p.sprite.setUniform("y", p.y)
}

func (p *pacman) draw() {
	p.sprite.draw()
}

func (p *pacman) onKey(key glfw.Key) {
	switch key {
	case glfw.KeyA:
		p.dir = 0
		p.x--
	case glfw.KeyW:
		p.dir = 1
		p.y++
	case glfw.KeyD:
		p.dir = 2
		p.x++
	case glfw.KeyS:
		p.dir = 3
		p.y--
	}
}

func newPacman() pacman {
	sprite := newSprite(
		"shaders/vert/pacman.glsl",
		"shaders/frag/pacman.glsl",
		"",
		-1, 1, 1./28.*3, 1./31.*3,
		[]string{
			"t",
			"dir",
			"x",
			"y",
		},
	)

	return pacman{sprite, 0, 0, 0}
}
