package main

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

type pacman struct {
	sprite *sprite
	dir    float64
	x      float64
	y      float64
}

func (p *pacman) update() {
	if isKeyPressed(glfw.KeyW) {
		p.dir = 3
		p.y += 8
	}
	if isKeyPressed(glfw.KeyA) {
		p.dir = 0
		p.x -= 8
	}
	if isKeyPressed(glfw.KeyS) {
		p.dir = 1
		p.y -= 8
	}
	if isKeyPressed(glfw.KeyD) {
		p.dir = 2
		p.x += 8
	}

	p.sprite.setUniform("t", float32(glfw.GetTime()))
	p.sprite.setUniform("dir", float32(p.dir))

	x, y := pxToScreen(p.x*2, p.y*2)

	p.sprite.setUniform("x", float32(x))
	p.sprite.setUniform("y", float32(y))
}

func (p *pacman) draw() {
	p.sprite.draw()
}

func newPacman() pacman {
	w, h := pxToScreen(16, 16)

	sprite := newSprite(
		"res/shaders/vert/pacman.glsl",
		"res/shaders/frag/pacman.glsl",
		"",
		-1, -1, float32(w), float32(h),
		[]string{
			"t",
			"dir",
			"x",
			"y",
		},
	)

	return pacman{sprite, 0, 4, 4}
}
