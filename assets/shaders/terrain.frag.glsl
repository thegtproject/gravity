#version 430 core

uniform sampler2D uSplatmap;
uniform sampler2D uHeightmap;
uniform sampler2D uRock;
uniform sampler2D uDirt;
uniform sampler2D uGrass;

in vec2 vCoord;

out vec4 fragColor;

void main() { 	
	vec4 splat = texture2D(uSplatmap, vCoord);
	vec4 height = texture2D(uHeightmap, vCoord);
	vec4 dirt  = texture2D(uDirt,  vCoord*128.0);
	vec4 rock  = texture2D(uRock,  vCoord*128.0);
	vec4 grass = texture2D(uGrass, vCoord*128.0);
	vec4 red   = vec4(1.0, 0.0, 0.0, 1.0);
	vec4 green = vec4(0.0, 1.0, 0.0, 1.0);
	vec4 blue  = vec4(0.0, 0.0, 1.0, 1.0);
	vec4 alpha = vec4(0.0, 0.0, 0.0, 1.0);

	fragColor = rock * splat.r;
	fragColor += mix(fragColor, dirt, splat.g);
	fragColor += mix(fragColor, rock, splat.b);
	fragColor += mix(fragColor, dirt, splat.a);
	fragColor = vec4(mix(grass*0.3,  fragColor, height.r-0.2).rgb, 1.0);


	// fragColor = red   * (0.0 - splat.r) +
	// 			green * (0.0 - splat.g) +
	// 			blue  * (0.0 - splat.b) +
	// 			alpha * (1.0 - splat.a);
}
