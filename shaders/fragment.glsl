#version 150 core

in vec2 uv;
out vec4 gl_Color;

void main() {
  gl_Color = vec4(uv, 0.0, 1.0);
}