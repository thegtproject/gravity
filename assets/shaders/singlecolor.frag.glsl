#version 410 core
 
in vec4 vDiffuse;

out vec4 fragColor;

void main() {
	fragColor = vDiffuse;
}
