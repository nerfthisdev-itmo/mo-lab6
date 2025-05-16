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

	newGen := make([]Genome, 0)

	for i := 0; i < len(population)/2; i++ {
		parent1 := selectParentByRoulette(population)
		parent2 := selectParentByRoulette(population)

		for parent1.ID == parent2.ID {
			parent2 = selectParentByRoulette(population)
		}

		child1, child2 := parent1.Reproduce(parent2)
		newGen = append(newGen, child1, child2)
	}

	population = append(population, newGen...)

	printOverview(population)
	population = reduce(population)

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
	fmt.Printf("%-5s | %-20s | %-5s | %-10s\n", "ID", "Chromosome", "Cost", "P(select)")
	fmt.Println("-------------------------------------------------------------")

	total := 0.0
	probabilities := make([]float64, len(population))

	for i, g := range population {
		if g.Cost == 0 {
			probabilities[i] = 1e6
		} else {
			probabilities[i] = 1.0 / float64(g.Cost)
		}
		total += probabilities[i]
	}

	for i := range probabilities {
		probabilities[i] /= total
	}

	for i, g := range population {
		fmt.Printf("%-5d | %-v | %-5d | %.6f\n", g.ID, g.Chromosome, g.Cost, probabilities[i])
	}

	fmt.Println()
}

func selectParentByRoulette(population []Genome) Genome {
	total := 0.0
	probabilities := make([]float64, len(population))

	for i, g := range population {
		if g.Cost == 0 {
			probabilities[i] = 1e6
		} else {
			probabilities[i] = 1.0 / float64(g.Cost)
		}
		total += probabilities[i]
	}

	for i := range probabilities {
		probabilities[i] /= total
	}

	r := rand.Float64()
	cumulative := 0.0
	for i, p := range probabilities {
		cumulative += p
		if r <= cumulative {
			return population[i]
		}
	}

	return population[len(population)-1]
}

func sortByCostDesc(population []Genome) {
	sort.Slice(population, func(i, j int) bool {
		return population[i].Cost > population[j].Cost
	})
}
