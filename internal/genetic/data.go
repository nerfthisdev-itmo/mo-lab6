package genetic

import "math/rand"

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Magenta = "\033[35m"

var Paths [][]int = [][]int{
	{0, 1, 7, 2, 8},  // City 1
	{2, 0, 10, 3, 1}, // City 2
	{7, 10, 0, 2, 6}, // City 3
	{2, 3, 2, 0, 4},  // City 4
	{8, 1, 6, 4, 0},  // City 5
}

var probability float64 = 0.01

var genomeIDCounter int

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
