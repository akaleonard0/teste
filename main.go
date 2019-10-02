package main

import (
	"app12/model"
	"fmt"
)

func main() {

	jojo := model.Ave{}
	jojo.Nome = "Jojooooooo da silva"

	queroAcordarComUmCarcarejo(jojo)
	queroOuvirUmaPAtaNoLago(jojo)
}

func queroAcordarComUmCarcarejo(g model.Galinha) {
	fmt.Println(g.Cacareja())
}

func queroOuvirUmaPAtaNoLago(p model.Pata) {
	fmt.Println(p.Quack())
}
