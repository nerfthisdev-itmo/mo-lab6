package main

import (
	"fmt"

	"github.com/nerfthisdev-itmo/mo-lab6/internal/genetic"
)

func main() {
	var n int

	fmt.Println("Please input the number of generations: ")
	fmt.Scan(&n)

	population := genetic.GeneratePopulation(4)

	for i := range n {
		fmt.Printf("GEN %d\n", i+1)
		fmt.Println("------------------------------------")
		population = genetic.Evolve(population)
		fmt.Println("------------------------------------")
	}
}
