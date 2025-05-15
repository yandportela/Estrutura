package main

import "fmt"

type Carro struct {
	modelo string
	ano int
	velocidade int
}



func exibirDados(c Carro) {
	fmt.Println("\nModelo:", c.modelo)
	fmt.Println("Ano:", c.ano)
	fmt.Println("Velocidade máxima:", c.velocidade, "\n")
}

func main() {
	c1 := Carro{"Onix", 2013, 167}
	c2 := Carro{"Tesla Cybertruck", 2023, 209}
	c3 := Carro{"Celta", 2006, 157}
	exibirDados(c1)
	exibirDados(c2)
	exibirDados(c3)
}