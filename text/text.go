package text

import (
	"context"
	"fmt"
	"math/rand"
	"time"
	"unicode"

	"github.com/caseylmanus/gophercon-ga/combinations"
	"github.com/caseylmanus/gophercon-ga/gen"
)

func Solve(targetString string, printUp func(string)) {
	target := []rune(targetString)
	validGenes := validRunes()
	species := gen.Species[rune]{
		ValidGenes:          validGenes,
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
	reporter := func(r gen.Report[rune]) {
		printUp(fmt.Sprintln("Species: ", r.Species, "Generation:", r.Generation, "Fitness:", r.Fittest.Fitness, "Value:", string(r.Fittest.Value)))
		if r.Fittest.Fitness == 1 {
			printUp(fmt.Sprintln("There were", combinations.Possible(int64(len(target)), int64(len(validGenes))).String(), "combinations!"))
		}
	}
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	species.Solve(context.Background(), random, 1, reporter)
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
