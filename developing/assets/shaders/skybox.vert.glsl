#version 430 core

in vec3 aPosition;

uniform mat4 uViewMatrix;
uniform mat4 uProjectionMatrix;

out vec3 vCoord;

void main() {
    vCoord = aPosition;
    gl_Position = uProjectionMatrix * uViewMatrix * vec4(aPosition, 1.0);
}  
