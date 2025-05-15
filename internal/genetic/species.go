package genetic

import "math/rand"

var Paths [][]int = [][]int{
	{0, 1, 7, 2, 8},  // City 1
	{2, 0, 10, 3, 1}, // City 2
	{7, 10, 0, 2, 6}, // City 3
	{2, 3, 2, 0, 4},  // City 4
	{8, 1, 6, 4, 0},  // City 5
}

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
	i, j := getRandomIndexes()
	g.Chromosome[i], g.Chromosome[j] = g.Chromosome[j], g.Chromosome[i]
}

func (parent1 Genome) Reproduce(parent2 Genome) (Genome, Genome) {
	br1, br2 := getRandomIndexes()

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

	return NewGenome(child1), NewGenome(child2)

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
		population[i] = createSpecies()
	}

	return population
}

func getRandomIndexes() (int, int) {
	length := 4
	a := rand.Intn(length)
	b := rand.Intn(length)

	for a == b {
		b = rand.Intn(length)
	}

	if a < b {
		return a, b
	}

	return b, a

}
