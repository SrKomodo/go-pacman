#version 150 core

uniform float t;

in vec2 uv;
out vec4 gl_Color;

void main() {
  float r = sqrt(pow(uv.x - .5, 2) + pow(uv.y - .5, 2));
  float a = atan(uv.y - .5, uv.x - .5);

  float sinT = sin(t * 10) / 2 + .5;
  float mask = min(
    step(r, .5),
    step(abs(a / 3.14159), 1 - (sinT * .25))
  );

  gl_Color = mix(
    vec4(0, 0, 1, 0),
    vec4(1, 1, 0, 1),
    mask
  );
}
