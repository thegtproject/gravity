#version 430 core

layout(location = 3) uniform vec3 uDiffuse;

out vec4 fragColor;

void main() {
	fragColor = vec4(uDiffuse, 1.0);
}
