package queens

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/caseylmanus/gophercon-ga/combinations"
	"github.com/caseylmanus/gophercon-ga/gen"
)

func Solve(boardSize int, concurrency int, printUp func(string)) {
	reportCh := make(chan gen.Report[Point])
	reporter := func(r gen.Report[Point]) {
		reportCh <- r
	}
	validGenes := getValidPoints(boardSize)
	species := gen.Species[Point]{
		ValidGenes:          validGenes,
		GenomeSize:          boardSize,
		PopulationSize:      10000,
		MutationRate:        0.01,
		SingleCrossOverRate: .8,
		Fitness:             fitness,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for i := 0; i < concurrency; i++ {
		random := rand.New(rand.NewSource(time.Now().UnixNano()))
		go species.Solve(ctx, random, i+1, reporter)
	}
	for {
		select {
		case <-ctx.Done():
			return
		case r := <-reportCh:
			printUp(fmt.Sprintln("Species: ", r.Species, "Generation:", r.Generation, "Fitness:", r.Fittest.Fitness, "Value:", r.Fittest.Value))
			if r.Fittest.Fitness == 1 {
				printUp(fmt.Sprintln("There were", combinations.Possible(int64(boardSize), int64(len(validGenes))).String(), "combinations!"))
				return
			}
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

func getValidPoints(boardSize int) []Point {
	var points []Point
	for x := 0; x < boardSize; x++ {
		for y := 0; y < boardSize; y++ {
			points = append(points, Point{x: x, y: y})
		}
	}
	return points
}
