#version 150 core

in vec2 coords;
in vec2 uv;

out vec2 p;

void main () {
  p = uv;
  gl_Position = vec4(coords, 0.0, 1.0);
}
