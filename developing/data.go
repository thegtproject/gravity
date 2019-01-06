package main

var vsBasicShader = `
#version 330 core

layout(location = 0) in vec3 aPosition;
layout(location = 1) in vec4 aColor;

uniform mat4 uModelMatrix;
uniform mat4 uViewMatrix;
uniform mat4 uProjectionMatrix;

out vec4 vColor;

void main() {
	gl_Position = uProjectionMatrix * uViewMatrix * uModelMatrix * vec4(aPosition, 1);
	vColor = aColor;
}
`

var fsBasicShader = `
#version 330 core

in vec4 vColor;
out vec4 fragColor;

void main() {
	fragColor = vColor;
}
`
