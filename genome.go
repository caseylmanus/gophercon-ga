package main

import (
	"math/rand"
	"strings"
)

type Genome struct {
	Value      string
	Fitness    float64
	Generation int
}

func (g *Genome) Random(random *rand.Rand, length int) {
	var sb strings.Builder
	for i := 0; i < length; i++ {
		index := rand.Intn(len(validGenes))
		sb.WriteByte(validGenes[index])
	}
	g.Value = sb.String()
}

func (g *Genome) Mutate(rand *rand.Rand) Genome {
	mutant := Genome{}
	geneIndex := rand.Intn(len(validGenes))
	parentIndex := rand.Intn(len(g.Value))
	if parentIndex > 0 {
		mutant.Value += g.Value[:parentIndex]
	}
	mutant.Value += validGenes[geneIndex : 1+geneIndex]
	if parentIndex+1 < len(g.Value) {
		mutant.Value += g.Value[parentIndex+1:]
	}
	mutant.Generation = g.Generation + 1
	return mutant
}

func (g *Genome) SetFitness(f func(Genome) float64) {
	g.Fitness = f(*g)
}

var validGenes = " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!.,0123456789"
