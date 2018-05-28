package main

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

type pacman struct {
	sprite *sprite
	dir    float64
	from   *node
}

func (p *pacman) goTo(n *node) {
	if n != nil {
		p.from = n
	}
}

func (p *pacman) update() {
	if isKeyPressed(glfw.KeyA) {
		p.dir = 0
		p.goTo(p.from.connections[3].node)
	}
	if isKeyPressed(glfw.KeyS) {
		p.dir = 3
		p.goTo(p.from.connections[2].node)
	}
	if isKeyPressed(glfw.KeyD) {
		p.dir = 2
		p.goTo(p.from.connections[1].node)
	}
	if isKeyPressed(glfw.KeyW) {
		p.dir = 1
		p.goTo(p.from.connections[0].node)
	}

	p.sprite.setUniform("t", float32(glfw.GetTime()))
	p.sprite.setUniform("dir", float32(p.dir))

	x, y := nodeToScreen(float64(p.from.x), float64(p.from.y))

	p.sprite.setUniform("x", float32(x))
	p.sprite.setUniform("y", float32(y))
}

func (p *pacman) draw() {
	p.sprite.draw()
}

func newPacman() pacman {
	w, h := pxToScreen(8, 8)

	sprite := newSprite(
		"res/shaders/vert/pacman.glsl",
		"res/shaders/frag/pacman.glsl",
		"",
		0, 0, float32(w), float32(h),
		[]string{
			"t",
			"dir",
			"x",
			"y",
		},
	)

	return pacman{sprite, 0, lvl.getNode(1, 1)}
}
