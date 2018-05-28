package main

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

func pxToScreen(x, y float64) (float64, float64) {
	return 1.0 / width * x * 2, 1.0 / height * y * 2
}

func nodeToScreen(x, y float64) (float64, float64) {
	x = x - 1
	y = float64(lvl.h-1) - y - 1
	offX, offY := pxToScreen(4, 4)
	return offX + (x/float64(lvl.w-1))*2.0 - 1.0, offY + (y/float64(lvl.h-1))*2.0 - 1.0
}

func isKeyPressed(key glfw.Key) bool {
	return glfw.GetCurrentContext().GetKey(key) == glfw.Press
}
