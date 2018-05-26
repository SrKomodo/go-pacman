#version 150 core

uniform sampler2D tex;

in vec2 p;
out vec4 gl_Color;

void main() {
  gl_Color = texture(tex, p);
}
