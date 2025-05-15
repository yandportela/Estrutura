package main

import "fmt"

type Jogador struct {
	nome string
	vida int
	nivel int
}

func exibirDados(j Jogador) {
	fmt.Println("Nome:", j.nome)
	fmt.Println("Vida:", j.vida)
	fmt.Println("Nivel:", j.nivel)
}

func main() {
	j1 := Jogador{"Sim", 100, 1}
	exibirDados(j1)
}