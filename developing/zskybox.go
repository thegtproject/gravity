package main

import (
	"github.com/thegtproject/gravity"
	"github.com/thegtproject/gravity/internal/gravitygl"
	"github.com/thegtproject/gravity/pkg/materials"
	"github.com/thegtproject/gravity/pkg/mesh"
)

func configureSkybox() {

	tex := gravitygl.NewCubeMap(
		"assets/skybox/elyhills/hills_rt.tga",
		"assets/skybox/elyhills/hills_lf.tga",
		"assets/skybox/elyhills/hills_ft.tga",
		"assets/skybox/elyhills/hills_bk.tga",
		"assets/skybox/elyhills/hills_up.tga",
		"assets/skybox/elyhills/hills_dn.tga",
	)

	// tex := gravitygl.NewCubeMap(
	// 	"assets/skybox/eezabad/eezabad_rt.jpg",
	// 	"assets/skybox/eezabad/eezabad_lf.jpg",
	// 	"assets/skybox/eezabad/eezabad_ft.jpg",
	// 	"assets/skybox/eezabad/eezabad_bk.jpg",
	// 	"assets/skybox/eezabad/eezabad_up.jpg",
	// 	"assets/skybox/eezabad/eezabad_dn.jpg",
	// )

	// tex := gravitygl.NewCubeMap(
	// 	"assets/skybox/ocean/right.jpg",
	// 	"assets/skybox/ocean/left.jpg",
	// 	"assets/skybox/ocean/front.jpg",
	// 	"assets/skybox/ocean/back.jpg",
	// 	"assets/skybox/ocean/top.jpg",
	// 	"assets/skybox/ocean/bottom.jpg",
	// )

	m := mesh.NewCubeMap()
	m.Scale(6000.0)
	skybox = gravity.NewModel(m, materials.NewSkyBox(), cam)
	skybox.AddUniform("uSkyBox", tex)
	// terrain.AddUniform("uSplatmap", splatmap)

}
