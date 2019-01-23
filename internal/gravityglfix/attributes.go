package gravitygl

// Attributes ...
type Attributes struct{ data map[string]Attribute }

// Uniforms ...
type Uniforms struct{ data map[string]Uniform }

// Attribute ...
type Attribute struct {
	Name string
	Loc  GLAttribute
	Type EnumType
	Size int
}

// Uniform ...
type Uniform struct {
	Name string
	Loc  GLUniform
	Type EnumType
	Size int
}

// Add ...
func (attr *Attributes) Add(Name string, Loc GLAttribute, Type EnumType, Size int) {
	attr.data[Name] = Attribute{
		Name: Name,
		Loc:  Loc,
		Type: Type,
		Size: Size,
	}
}

// Get ...
func (attr *Attributes) Get(Name string) (val Attribute, ok bool) {
	val, ok = attr.data[Name]
	return
}

// Add ...
func (unif *Uniforms) Add(Name string, Loc GLUniform, Type EnumType, Size int) {
	unif.data[Name] = Uniform{
		Name: Name,
		Loc:  Loc,
		Type: Type,
		Size: Size,
	}
}

// Get ...
func (unif *Uniforms) Get(Name string) (val Uniform, ok bool) {
	val, ok = unif.data[Name]
	return
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
