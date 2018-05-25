#version 150 core

in vec2 p;
in vec2 _uv;

out vec2 uv;

void main () {
  uv = _uv;
  gl_Position = vec4(p, 0.0, 1.0);
}
