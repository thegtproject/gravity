#version 410 core

layout(location = 0) in vec3 aPosition;
layout(location = 1) in vec3 aDiffuse;

uniform mat4 uModelMatrix;
uniform mat4 uViewMatrix;
uniform mat4 uProjectionMatrix;

layout(location = 0) out vec4 vColor;

void main() {
	gl_Position = uProjectionMatrix * uViewMatrix * uModelMatrix * vec4(aPosition, 1);
	vColor = vec4(aDiffuse, 1.0);
}