#version 430 core

layout(location = 0) in vec3 aPosition;
layout(location = 2) in vec2 aCoord;

uniform mat4 uModelMatrix;
uniform mat4 uViewMatrix;
uniform mat4 uProjectionMatrix;

out vec2 vCoord;
out vec4 vColor;

void main() {
	gl_Position = uProjectionMatrix * uViewMatrix * uModelMatrix * vec4(aPosition, 1);
    vCoord = aCoord;
}
