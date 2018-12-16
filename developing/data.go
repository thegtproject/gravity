package main

var cubeindices = []uint16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35}

var cubevertices = []float32{
	-0.5, -0.5, -0.5,
	-0.5, -0.5, 0.5,
	-0.5, 0.5, 0.5,
	0.5, 0.5, -0.5,
	-0.5, -0.5, -0.5,
	-0.5, 0.5, -0.5,
	0.5, -0.5, 0.5,
	-0.5, -0.5, -0.5,
	0.5, -0.5, -0.5,
	0.5, 0.5, -0.5,
	0.5, -0.5, -0.5,
	-0.5, -0.5, -0.5,
	-0.5, -0.5, -0.5,
	-0.5, 0.5, 0.5,
	-0.5, 0.5, -0.5,
	0.5, -0.5, 0.5,
	-0.5, -0.5, 0.5,
	-0.5, -0.5, -0.5,
	-0.5, 0.5, 0.5,
	-0.5, -0.5, 0.5,
	0.5, -0.5, 0.5,
	0.5, 0.5, 0.5,
	0.5, -0.5, -0.5,
	0.5, 0.5, -0.5,
	0.5, -0.5, -0.5,
	0.5, 0.5, 0.5,
	0.5, -0.5, 0.5,
	0.5, 0.5, 0.5,
	0.5, 0.5, -0.5,
	-0.5, 0.5, -0.5,
	0.5, 0.5, 0.5,
	-0.5, 0.5, -0.5,
	-0.5, 0.5, 0.5,
	0.5, 0.5, 0.5,
	-0.5, 0.5, 0.5,
	0.5, -0.5, 0.5,
}

var vsNormalShading = `
#version 330 core

in vec3 aPosition;

uniform mat4 uModelMatrix;
uniform mat4 uViewMatrix;
uniform mat4 uProjectionMatrix;

out vec3 world_pos;

void main() {
	gl_Position = uProjectionMatrix * uViewMatrix * uModelMatrix * vec4(aPosition, 1);
	world_pos = gl_Position.xyz;
}
`

var fsNormalShading = `
#version 330 core
precision mediump float;
in vec3 world_pos;
out vec4 fragColor;
float checkersGradBox( in vec2 p ) {
    vec2 w = fwidth(p) + 0.001;
    vec2 i = 2.0*(abs(fract((p-0.5*w)*0.5)-0.5)-abs(fract((p+0.5*w)*0.5)-0.5))/w;
    return 0.5 - 0.5*i.x*i.y;                  
}

void main() {
	vec3 dFdxPos = dFdx( world_pos );
	vec3 dFdyPos = dFdy( world_pos );
	vec3 facenormal = normalize( cross(dFdxPos,dFdyPos ));
	fragColor = vec4(facenormal*0.5+0.5, 0.5);
}
`
