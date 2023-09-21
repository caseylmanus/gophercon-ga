package greeting

import (
	"fmt"
	"math/rand"
	"time"
	"unicode"

	"github.com/caseylmanus/gophercon-ga/combinations"
	"github.com/caseylmanus/gophercon-ga/gen"
)

func Solve(printUp func(string)) {
	target := []rune("Hello Gophercon 2023, Welcome to San Diego!")
	species := gen.Species[rune]{
		ValidGenes:          validRunes(),
		GenomeSize:          len(target),
		PopulationSize:      5000,
		MutationRate:        0.01,
		SingleCrossOverRate: .8,
		Fitness: func(g *gen.Genome[rune]) float64 {
			matches := 0
			for i := 0; i < len(target); i++ {
				if g.Value[i] == target[i] {
					matches++
				}
			}
			return float64(matches) / float64(len(target))
		},
	}
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	pop := species.RandomPopulation(random)
	bestScore := pop.Fittest.Fitness
	for {
		pop = species.NextPopulation(random, pop)
		if pop.Fittest.Fitness == 1 {
			printUp(fmt.Sprintln("Generation:", pop.Generation, "Fitness:", pop.Fittest.Fitness, "Value:", string(pop.Fittest.Value)))
			printUp(fmt.Sprintln("There were", combinations.Possible(int64(species.GenomeSize), int64(len(species.ValidGenes))).String(), "combinations!"))
			return
		}
		if pop.Fittest.Fitness > bestScore {
			bestScore = pop.Fittest.Fitness
			printUp(fmt.Sprintln("Generation:", pop.Generation, "Fitness:", pop.Fittest.Fitness, "Value:", string(pop.Fittest.Value)))
		}
	}
}

func validRunes() []rune {
	var results []rune
	for i := 0; i < unicode.MaxASCII; i++ {
		if unicode.IsPrint(rune(i)) {
			results = append(results, rune(i))
		}
	}
	return results
}
