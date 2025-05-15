package genetic

import "math/rand"

var Paths [][]int = [][]int{
	{0, 1, 7, 2, 8},  // City 1
	{2, 0, 10, 3, 1}, // City 2
	{7, 10, 0, 2, 6}, // City 3
	{2, 3, 2, 0, 4},  // City 4
	{8, 1, 6, 4, 0},  // City 5
}

type Genome struct {
	Chromosome []int
	Cost       int
}

func (g *Genome) CreateSpecies() {
	numbers := []int{0, 1, 2, 3, 4}

	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})

	g.Chromosome = numbers

	for i := range len(g.Chromosome) - 1 {
		from := g.Chromosome[i]
		to := g.Chromosome[i+1]
		g.Cost += Paths[from][to]
	}

	// return to 1st city in chain
	start := g.Chromosome[0]
	end := g.Chromosome[len(g.Chromosome)-1]
	g.Cost += Paths[end][start]

}
