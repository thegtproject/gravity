package opengl

// Attributes ...
type Attributes struct{ data map[string]Attribute }

// Uniforms ...
type Uniforms struct{ data map[string]Uniform }

// Attribute ...
type Attribute struct {
	name string
	loc  AttributeLoc
	ty   Enum
	size int
}

// Uniform ...
type Uniform struct {
	name string
	loc  UniformLoc
	ty   Enum
	size int
}

// Add ...
func (attr *Attributes) Add(name string, loc AttributeLoc, ty Enum, size int) {
	attr.data[name] = Attribute{
		name: name,
		loc:  loc,
		ty:   ty,
		size: size,
	}
}

// Add ...
func (unif *Uniforms) Add(name string, loc UniformLoc, ty Enum, size int) {
	unif.data[name] = Uniform{
		name: name,
		loc:  loc,
		ty:   ty,
		size: size,
	}
}

// NewAttributesMap ...
func NewAttributesMap() *Attributes {
	return &Attributes{
		data: make(map[string]Attribute),
	}
}

// NewUniformsMap ...
func NewUniformsMap() *Uniforms {
	return &Uniforms{
		data: make(map[string]Uniform),
	}
}
