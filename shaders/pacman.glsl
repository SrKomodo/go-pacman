#version 150 core

in vec2 uv;
out vec4 gl_Color;

void main() {
  float mask = step(distance(uv, vec2(0.5, 0.5)), 0.5);
  gl_Color = mix(
    vec4(0.0, 0.0, 1.0, 0.0),
    vec4(1.0, 1.0, 0.0, 1.0),
    mask
  );
}
