package main

import (
	"github.com/thegtproject/gravity"
	"github.com/thegtproject/gravity/internal/gravitygl"
	"github.com/thegtproject/gravity/pkg/materials"
	"github.com/thegtproject/gravity/pkg/mesh"
)

func configureSkybox() {
	tex := gravity.NewCubeMap(
		"assets/skybox/yukongold/yukongold_rt.tga",
		"assets/skybox/yukongold/yukongold_lf.tga",
		"assets/skybox/yukongold/yukongold_ft.tga",
		"assets/skybox/yukongold/yukongold_bk.tga",
		"assets/skybox/yukongold/yukongold_up.tga",
		"assets/skybox/yukongold/yukongold_dn.tga",
	)

	m := mesh.NewCubeMap()
	m.Scale(6000.0)
	skybox = gravity.NewModel(m, materials.NewSkyBox(), cam)
	skybox.AddUniform("uSkyBox", tex)
}

func configureLinewidget() {

	linewidget = gravity.NewModel(
		mesh.FromGob("assets/mesh/linewidget.gmesh"),
		materials.NewNone(), cam,
	)
	linewidget.Scalef(65)
	linewidget.TranslateZ(450)
	linewidget.Primitive = gravitygl.LINES

}

func configureTerrain() {
	splatmap := gravity.NewTextureFromFile(gravitygl.TEXTURE_2D, "assets/terrain/splatmap.png")
	heightmap := gravity.NewTextureFromFile(gravitygl.TEXTURE_2D, "assets/terrain/height.png")
	dirt := gravity.NewTextureFromFile(gravitygl.TEXTURE_2D, "assets/textures/dirt05.jpg")
	rock := gravity.NewTextureFromFile(gravitygl.TEXTURE_2D, "assets/textures/rock.jpg")
	grass := gravity.NewTextureFromFile(gravitygl.TEXTURE_2D, "assets/textures/grass.jpg")

	splatmap.Unit = 0
	heightmap.Unit = 1
	dirt.Unit = 2
	rock.Unit = 3
	grass.Unit = 4

	gravity.UploadToGPU(splatmap)
	gravity.UploadToGPU(heightmap)
	gravity.UploadToGPU(dirt)
	gravity.UploadToGPU(rock)
	gravity.UploadToGPU(grass)

	terrain = gravity.NewModel(mesh.FromGob("assets/mesh/terrain.obj"), materials.NewTerrain(), cam)

	terrain.AddUniform("uSplatmap", splatmap)
	terrain.AddUniform("uHeightmap", heightmap)
	terrain.AddUniform("uDirt", dirt)
	terrain.AddUniform("uRock", rock)
	terrain.AddUniform("uGrass", grass)

}
