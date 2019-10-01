package model

//Galinha representa uma ave do tipo galinha
type Galinha interface {
	Cacareja() string
}

//Pata representa uma ave
type Pata interface {
	Quack() string
}

//Ave representa um animal
type Ave struct {
	Nome string
}

//Cacareja representa o som de uma galinha
func (a Ave) Cacareja() string {
	return "Cócóricóooooooo..."
}

// Quack representa o som de um pato
func (a Ave) Quack() string {
	return "Quack, quackkkkkkkk..."
}
