package gravitygl

// Program ...
type Program struct {
	id uint32
}

// NewProgram ...
func NewProgram(vertexSource, fragmentSource string) *Program {
	prg, err := MakeProgram(vertexSource, fragmentSource)
	if err != nil {
		panic(err)
	}
	return &Program{
		id: prg,
	}
}

// Use ...
func (prg *Program) Use() {
	UseProgram(prg.id)
}
