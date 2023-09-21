package queens

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/caseylmanus/gophercon-ga/combinations"
	"github.com/caseylmanus/gophercon-ga/gen"
)

func Solve(printUp func(string)) {
	species := gen.Species[Point]{
		ValidGenes:          getValidPoints(),
		GenomeSize:          8,
		PopulationSize:      10000,
		MutationRate:        0.01,
		SingleCrossOverRate: .8,
		Fitness:             fitness,
	}
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	pop := species.RandomPopulation(random)
	bestScore := pop.Fittest.Fitness
	for {
		pop = species.NextPopulation(random, pop)
		if pop.Fittest.Fitness == 1 {
			printUp(fmt.Sprintln("Generation:", pop.Generation, "Fitness:", pop.Fittest.Fitness, "Value:", pop.Fittest.Value))
			printUp(fmt.Sprintln("There were", combinations.Possible(int64(species.GenomeSize), int64(len(species.ValidGenes))).String(), "combinations!"))
			return
		}
		if pop.Fittest.Fitness > bestScore {
			bestScore = pop.Fittest.Fitness
			printUp(fmt.Sprintln("Generation:", pop.Generation, "Fitness:", pop.Fittest.Fitness, "Value:", pop.Fittest.Value))
		}
	}
}

type Point struct {
	x, y int
}

func fitness(genome *gen.Genome[Point]) float64 {
	score := 0
	for i, g := range genome.Value {
		existing := genome.Value[0:i]
		if CanPlace(g, existing) {
			score++
		}
	}
	return float64(score) / float64(len(genome.Value))
}

func CanPlace(target Point, board []Point) bool {
	for _, point := range board {
		if CanAttack(point, target) {
			return false
		}
	}
	return true
}

func CanAttack(a, b Point) bool {
	answer := a.x == b.x || a.y == b.y || math.Abs(float64(a.y-b.y)) == math.Abs(float64(a.x-b.x))
	return answer
}

func getValidPoints() []Point {
	var points []Point
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			points = append(points, Point{x: x, y: y})
		}
	}
	return points
}
