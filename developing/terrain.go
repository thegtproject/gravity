package main

import (
	"github.com/thegtproject/gravity"
	"github.com/thegtproject/gravity/internal/gravitygl"
	"github.com/thegtproject/gravity/pkg/materials"
	"github.com/thegtproject/gravity/pkg/mesh"
)

func configureTerrain() {
	splatmap := gravitygl.NewTextureFromFile("assets/terr/splatmap.png")
	heightmap := gravitygl.NewTextureFromFile("assets/terr/height.png")
	dirt := gravitygl.NewTextureFromFile("../assets/textures/terrain/dirt05.jpg")
	rock := gravitygl.NewTextureFromFile("../assets/textures/terrain/rock.jpg")
	grass := gravitygl.NewTextureFromFile("../assets/textures/terrain/grass.jpg")

	splatmap.Unit = 0
	heightmap.Unit = 1
	dirt.Unit = 2
	rock.Unit = 3
	grass.Unit = 4

	gravitygl.UploadToGPU(splatmap)
	gravitygl.UploadToGPU(heightmap)
	gravitygl.UploadToGPU(dirt)
	gravitygl.UploadToGPU(rock)
	gravitygl.UploadToGPU(grass)

	terrain = gravity.NewModel(mesh.FromGob("assets/terr/terrain.obj"), materials.NewTerrain(), cam)

	terrain.AddUniform("uSplatmap", splatmap)
	terrain.AddUniform("uHeightmap", heightmap)
	terrain.AddUniform("uDirt", dirt)
	terrain.AddUniform("uRock", rock)
	terrain.AddUniform("uGrass", grass)

}
