package main

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

func pxToScreen(x, y float64) (float64, float64) {
	return 1.0 / width * x, 1.0 / height * y
}

func isKeyPressed(key glfw.Key) bool {
	return glfw.GetCurrentContext().GetKey(key) == glfw.Press
}
