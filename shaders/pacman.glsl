#version 150 core

uniform float t;
uniform float dir;

in vec2 p;
out vec4 gl_Color;

void main() {
  float r = sqrt(pow(p.x - .5, 2) + pow(p.y - .5, 2));
  float a = atan(p.y - .5, p.x - .5);

  float sinT = sin(t * 10) / 2 + .5;
  float mask = min(
    step(r, .5),
    step(
      abs(
        mod(
          a / 3.14159  // [-1;1]
          / 2 + .5     // [ 0;1]
          - dir / 4, 1 // [ 0;1] centered at dir
        ) * 2 - 1      // [-1;1] centered at dir
      ),
      1 - (sinT * .25)
    )
  );

  gl_Color = mix(
    vec4(0, 0, 1, 0),
    vec4(1, 1, 0, 1),
    mask
  );
}
