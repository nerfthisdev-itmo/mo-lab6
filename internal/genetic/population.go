package genetic

import (
	"fmt"
	"math/rand"
	"sort"
)

func GeneratePopulation(n int) []Genome {
	population := make([]Genome, n)

	for i := range n {
		g := createSpecies()
		population[i] = g
		fmt.Printf("Generated new species ID: %d Chromosome: %v Cost: %d\n", g.ID, g.Chromosome, g.Cost)
	}

	printOverview(population)

	return population
}

func Evolve(population []Genome) []Genome {
	fmt.Println(Green + "Starting the evolution cycle..." + Reset)

	indexes := rand.Perm(len(population))

	for i := 0; i < len(indexes)-1; i += 2 {
		parent1 := population[indexes[i]]
		parent2 := population[indexes[i+1]]

		child1, child2 := parent1.Reproduce(parent2)

		population = append(population, child1, child2)
	}

	printOverview(population)

	population = reduce(population) // ← ВАЖНО: обновляем переменную

	fmt.Println(Green + "Evolution cycle has ended" + Reset)

	return population
}

func reduce(population []Genome) []Genome {
	sortByCostDesc(population)
	for i := 0; i <= len(population)/2-1; i++ {
		g := population[i]
		fmt.Printf(Red+"Killed species ID: %d Chromosome: %v Cost: %d\n"+Reset, g.ID, g.Chromosome, g.Cost)
	}

	newPopulation := population[len(population)/2:]

	printOverview(newPopulation)
	return newPopulation
}

func printOverview(population []Genome) {
	fmt.Println()
	fmt.Printf("%-5s | %-5s | %-5s\n", "ID", "Chromosome", "Cost")
	fmt.Println("-----------------------------")
	for _, g := range population {
		fmt.Printf("%-5d | %v | %-5d\n", g.ID, g.Chromosome, g.Cost)
	}
	fmt.Println()
}

func sortByCostDesc(population []Genome) {
	sort.Slice(population, func(i, j int) bool {
		return population[i].Cost > population[j].Cost
	})
}
