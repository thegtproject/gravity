package main

import (
	"github.com/thegtproject/gravity"
	"github.com/thegtproject/gravity/internal/gravitygl"
	"github.com/thegtproject/gravity/pkg/materials"
	"github.com/thegtproject/gravity/pkg/mesh"
)

func configureLinewidget() {

	linewidget = gravity.NewModel(
		mesh.FromGob("assets/linewidget.gmesh"),
		materials.NewNone(), cam,
	)
	linewidget.Scalef(35)
	linewidget.Primitive = gravitygl.LINES

}
