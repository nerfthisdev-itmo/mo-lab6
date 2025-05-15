package genetic

import (
	"fmt"
	"math/rand"
)

var Paths [][]int = [][]int{
	{0, 1, 7, 2, 8},  // City 1
	{2, 0, 10, 3, 1}, // City 2
	{7, 10, 0, 2, 6}, // City 3
	{2, 3, 2, 0, 4},  // City 4
	{8, 1, 6, 4, 0},  // City 5
}

var probability float64 = 0.01

var genomeIDCounter int

type Genome struct {
	ID         int
	Chromosome []int
	Cost       int
}

func (g *Genome) evaluate() {
	g.Cost = 0
	for i := range len(g.Chromosome) - 1 {
		from := g.Chromosome[i]
		to := g.Chromosome[i+1]
		g.Cost += Paths[from][to]
	}
	g.Cost += Paths[g.Chromosome[len(g.Chromosome)-1]][g.Chromosome[0]]
}

func NewGenome(chromosome []int) Genome {
	g := Genome{
		ID:         genomeIDCounter,
		Chromosome: chromosome,
	}
	genomeIDCounter++
	g.evaluate()
	return g
}

func (g *Genome) mutate() {
	i, j := getRandomIndexes(4)
	oldCost := g.Cost
	g.Chromosome[i], g.Chromosome[j] = g.Chromosome[j], g.Chromosome[i]
	g.evaluate()
	fmt.Printf("Species ID: %d has mutated!\n Cost changed from %d to %d", g.ID, oldCost, g.Cost)
}

func (parent1 Genome) Reproduce(parent2 Genome) (Genome, Genome) {
	br1, br2 := getRandomIndexes(4)

	segment1 := parent1.Chromosome[br1 : br2+1]
	segment2 := parent2.Chromosome[br1 : br2+1]

	child1 := make([]int, len(parent1.Chromosome))
	child2 := make([]int, len(parent2.Chromosome))

	for i := range child1 {
		child1[i] = -1
		child2[i] = -1
	}

	copy(child1[br1:br2+1], segment1)
	copy(child2[br1:br2+1], segment2)

	fillRemaining := func(child, segment, donor []int) {
		used := make(map[int]bool)
		for _, v := range segment {
			used[v] = true
		}

		startIdx := (br1 + 1) % len(donor)
		insertIdx := 0

		for i := range donor {
			gene := donor[(startIdx+i)%len(donor)]
			if !used[gene] {

				for child[insertIdx] != -1 {
					insertIdx++
				}
				child[insertIdx] = gene
				used[gene] = true
			}
		}
	}

	fillRemaining(child1, segment1, parent2.Chromosome)
	fillRemaining(child2, segment2, parent1.Chromosome)

	g1 := NewGenome(child1)
	g2 := NewGenome(child2)

	if rand.Float64() < probability {
		g1.mutate()
	}

	if rand.Float64() < probability {
		g2.mutate()
	}

	fmt.Printf("New child of parents %d and %d: %d Chromosome: %v Cost: %d\n", parent1.ID, parent2.ID, g1.ID, g1.Chromosome, g1.Cost)
	fmt.Printf("New child of parents %d and %d: %d Chromosome: %v Cost: %d\n", parent1.ID, parent2.ID, g2.ID, g2.Chromosome, g1.Cost)

	return g1, g2

}

func createSpecies() Genome {
	numbers := []int{0, 1, 2, 3, 4}
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})
	return NewGenome(numbers)
}

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

func Evolve(population []Genome) {
	fmt.Println("Starting the evolution cycle...")

	indexes := rand.Perm(len(population))

	for i := 0; i < len(indexes)-1; i += 2 {
		parent1 := population[indexes[i]]
		parent2 := population[indexes[i+1]]

		child1, child2 := parent1.Reproduce(parent2)

		population = append(population, child1, child2)
	}

	printOverview(population)

	fmt.Println("Evolution cycle has ended")
}

func getRandomIndexes(n int) (int, int) {
	a := rand.Intn(n)
	b := rand.Intn(n)

	for a == b {
		b = rand.Intn(n)
	}

	if a < b {
		return a, b
	}

	return b, a

}

func printOverview(population []Genome) {
	fmt.Printf("%-5s | %-5s | %-5s\n", "ID", "Chromosome", "Cost")
	fmt.Println("-----------------------------")
	for _, g := range population {
		fmt.Printf("%-5d | %v | %-5d\n", g.ID, g.Chromosome, g.Cost)
	}
}
