package main

import (
	"github.com/thegtproject/gravity"
	"github.com/thegtproject/gravity/materials"
	"github.com/thegtproject/gravity/mesh"
)

func configureTerrain() {
	splatmap := gravity.NewTextureFromFile("assets/terr/splatmap.png")
	heightmap := gravity.NewTextureFromFile("assets/terr/height.png")
	dirt := gravity.NewTextureFromFile("../assets/textures/terrain/dirt05.jpg")
	rock := gravity.NewTextureFromFile("../assets/textures/terrain/rock.jpg")
	grass := gravity.NewTextureFromFile("../assets/textures/terrain/grass.jpg")

	splatmap.Unit = 0
	heightmap.Unit = 1
	dirt.Unit = 2
	rock.Unit = 3
	grass.Unit = 4

	splatmap.UploadToGPU()
	heightmap.UploadToGPU()
	dirt.UploadToGPU()
	rock.UploadToGPU()
	grass.UploadToGPU()

	terrain = gravity.NewModel(mesh.FromGob("assets/terr/terrain.obj"), materials.NewTerrain(), cam)

	terrain.AddUniform("uSplatmap", splatmap)
	terrain.AddUniform("uHeightmap", heightmap)
	terrain.AddUniform("uDirt", dirt)
	terrain.AddUniform("uRock", rock)
	terrain.AddUniform("uGrass", grass)

}
