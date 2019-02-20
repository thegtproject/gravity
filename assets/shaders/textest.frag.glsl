#version 430 core

layout(location = 5) uniform sampler2D uTexture;

in vec4 vColor;
in vec2 vCoord;

out vec4 fragColor;

void main() {	 
 	fragColor = texture2D(uTexture, vCoord);
}
