package main

import (
	"app12/model"
	"fmt"
)

func main() {

	animal := model.Ave{}
	animal.Nome = "Jojo da silva"

	queroAcordarComUmCarcarejo(animal)
	queroOuvirUmaPAtaNoLago(animal)
}

func queroAcordarComUmCarcarejo(g model.Galinha) {
	fmt.Println(g.Cacareja())
}

func queroOuvirUmaPAtaNoLago(p model.Pata) {
	fmt.Println(p.Quack())
}
