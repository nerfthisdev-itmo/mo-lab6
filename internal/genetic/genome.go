package genetic

import (
	"fmt"
	"math/rand"
)

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

func (g *Genome) mutate() {
	i, j := getRandomIndexes(4)
	oldCost := g.Cost
	g.Chromosome[i], g.Chromosome[j] = g.Chromosome[j], g.Chromosome[i]
	g.evaluate()
	fmt.Printf(Magenta+"Species ID: %d has mutated! Cost changed from %d to %d\n"+Reset, g.ID, oldCost, g.Cost)
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
	fmt.Printf("New child of parents %d and %d: %d Chromosome: %v Cost: %d\n\n", parent1.ID, parent2.ID, g2.ID, g2.Chromosome, g1.Cost)

	return g1, g2

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

func createSpecies() Genome {
	numbers := []int{0, 1, 2, 3, 4}
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})
	return NewGenome(numbers)
}
