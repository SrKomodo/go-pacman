package main

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

type pacman struct {
	sprite  *sprite
	from    *node
	to      *connection
	travel  float64
	dir     int
	nextDir int
}

func (p *pacman) update(t, dT float64) {
	p.sprite.setUniform("t", float32(t))

	switch {
	case isKeyPressed(glfw.KeyW):
		p.nextDir = 0
	case isKeyPressed(glfw.KeyD):
		p.nextDir = 1
	case isKeyPressed(glfw.KeyS):
		p.nextDir = 2
	case isKeyPressed(glfw.KeyA):
		p.nextDir = 3
	}

	// If we want to reverse
	if (p.dir-p.nextDir == 2 || p.dir-p.nextDir == -2) && p.travel != 0 {
		p.from = p.to.node
		p.to = p.from.connections[p.nextDir]
		p.travel = 1 - p.travel
		p.dir = p.nextDir
		p.sprite.setUniform("dir", float32(p.dir))
	}

	if p.to == nil {
		if p.from.connections[p.nextDir] != nil {
			p.sprite.setUniform("dir", float32(p.nextDir))
			p.to = p.from.connections[p.nextDir]
			p.dir = p.nextDir
		}
	} else {
		p.travel += dT / float64(p.to.distance) * 4

		x := float64(p.from.x) + p.travel*float64(p.to.node.x-p.from.x)
		y := float64(p.from.y) + p.travel*float64(p.to.node.y-p.from.y)
		x, y = nodeToScreen(x, y)

		p.sprite.setUniform("x", float32(x))
		p.sprite.setUniform("y", float32(y))
	}

	if p.travel >= 1 {
		p.from = p.to.node
		p.travel = 0
		if p.to.node.connections[p.nextDir] != nil {
			p.sprite.setUniform("dir", float32(p.nextDir))
			p.to = p.from.connections[p.nextDir]
			p.dir = p.nextDir
		} else {
			if p.to.node.connections[p.dir] != nil {
				p.sprite.setUniform("dir", float32(p.dir))
				p.to = p.from.connections[p.dir]
			} else {
				p.to = nil
			}
		}
	}
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

	sprite.setUniform("dir", 1)
	return pacman{sprite, lvl.getNode(1, 1), lvl.getNode(1, 1).connections[1], 0, 1, 1}
}
