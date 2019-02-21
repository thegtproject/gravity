#version 430 core

in vec3 vCoord;

uniform samplerCube uSkyBox;

out vec4 fragColor;

void main() {
   fragColor = texture(uSkyBox, vCoord);
}