#version 150 core

uniform float x;
uniform float y;

in vec2 coords;
in vec2 uv;

out vec2 p;

void main () {
  p = uv;
  gl_Position = vec4(coords + vec2(x, y) / 100, 0.0, 1.0);
}
