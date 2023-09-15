package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	target := []byte("Hello Gophercon")
	populationSize := 1000
	mutationRate := 0.01

	population := make([][]byte, populationSize)
	for i := range population {
		population[i] = make([]byte, len(target))
		for j := range population[i] {
			population[i][j] = validGenes[rand.Intn(len(validGenes))]
		}
	}

	for generation := 0; ; generation++ {
		var best []byte
		var bestFitness int
		for _, individual := range population {
			fitness := 0
			for i := range individual {
				if individual[i] == target[i] {
					fitness++
				}
			}
			if fitness > bestFitness {
				best = individual
				bestFitness = fitness
			}
			if fitness == len(target) {
				fmt.Printf("Generation %d: %s\n", generation, string(individual))
				return
			}
		}

		fmt.Printf("Generation %d: %s\n", generation, string(best))

		nextPopulation := make([][]byte, populationSize)
		for i := range nextPopulation {
			parent1 := population[rand.Intn(len(population))]
			parent2 := population[rand.Intn(len(population))]
			child := make([]byte, len(target))
			for j := range child {
				if rand.Float64() < mutationRate {
					child[j] = validGenes[rand.Intn(len(validGenes))]
				} else if rand.Float64() < 0.5 {
					child[j] = parent1[j]
				} else {
					child[j] = parent2[j]
				}
			}
			nextPopulation[i] = child
		}

		population = nextPopulation
	}
}

var validGenes = makeValidGenes()

func makeValidGenes() []byte {
	allStrings := []string{" ", "!", ".", ","}
	for i := 0; i < 10; i++ {
		allStrings = append(allStrings, fmt.Sprintf("%d", i))
	}
	for c := 'a'; c <= 'z'; c++ {
		allStrings = append(allStrings, string(c))
	}
	for c := 'A'; c <= 'Z'; c++ {
		allStrings = append(allStrings, string(c))
	}
	var results []byte
	for _, s := range allStrings {
		results = append(results, []byte(s)...)
	}
	return results
}
