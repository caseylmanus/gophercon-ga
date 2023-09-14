package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	target := "Hello Gophercon 2023, Welcome to San Diego!"
	fitness := func(genome Genome) float64 {
		return getFitness(target, genome.Value)
	}

	start := time.Now()

	disp := func(g Genome, generation int) {
		fmt.Println("Generation:", generation, "Value:", g.Value, "Best Fitness: ", g.Fitness, "time", time.Since(start))
	}
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	disp(solve(fitness, disp, len(target), random))
	//disp(best, 0)
}

func solve(getFitness func(Genome) float64, display func(Genome, int), length int, random *rand.Rand) (Genome, int) {
	var bestParent Genome
	bestParent.Random(random, length)
	bestParent.SetFitness(getFitness)
	gen := 0
	for bestParent.Fitness < 1 {
		gen++
		child := bestParent.Mutate(random)
		child.SetFitness(getFitness)
		if child.Fitness > bestParent.Fitness {
			bestParent = child
		}
		display(bestParent, gen)
	}
	return bestParent, gen
}

func getFitness(target, candidate string) float64 {
	differenceCount := 0
	for i := 0; i < len(target); i++ {
		if target[i] != candidate[i] {
			differenceCount++
		}
	}
	return float64(len(target)-differenceCount) / float64(len(target))
}