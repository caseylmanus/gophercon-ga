package gen

import "fmt"

type FitnessFunc[T any] func(*Genome[T]) float64
type PrinterFunc[T any] func(string)

type Species[T any] struct {
	ValidGenes          []T
	Fitness             FitnessFunc[T]
	PopulationSize      int
	GenomeSize          int
	MutationRate        float64
	SingleCrossOverRate float64
	Printer             PrinterFunc[T]
	//other configuration
}

type Randomizer interface {
	Intn(i int) int
	Float64() float64
}

func (e *Species[T]) Solve(random Randomizer) {
	pop := e.RandomPopulation(random)
	bestScore := pop.Fittest.Fitness
	for {
		pop = e.NextPopulation(random, pop)
		if pop.Fittest.Fitness == 1 {
			fmt.Sprintln("Generation:", pop.Generation, "Fitness:", pop.Fittest.Fitness, "Value:", pop.Fittest.Value)
			//fmt.Println("There were", combinations.Possible(int64(species.GenomeSize), int64(len(species.ValidGenes))).String(), "combinations!")
			return
		}
		if pop.Fittest.Fitness > bestScore {
			bestScore = pop.Fittest.Fitness
			fmt.Println("Generation:", pop.Generation, "Fitness:", pop.Fittest.Fitness, "Value:", pop.Fittest.Value)
		}
	}
}

func (e *Species[T]) NextPopulation(random Randomizer, pop Population[T]) Population[T] {
	var next Population[T]
	next.Generation = pop.Generation + 1
	next.Genomes = make([]*Genome[T], e.PopulationSize)
	breedingPool := e.Selection(random, pop)
	poolSize := len(breedingPool)
	for i := 0; i < e.PopulationSize; i++ {
		a, b := breedingPool[random.Intn(poolSize)], breedingPool[random.Intn(poolSize)]
		genome := e.CrossOver(random, a, b)
		next.Genomes[i] = genome
	}
	next.SetFitness(e.Fitness)
	return next
}

func (e *Species[T]) RandomPopulation(random Randomizer) Population[T] {
	var population Population[T]
	population.Generation = 0
	population.Genomes = make([]*Genome[T], e.PopulationSize)
	for i := 0; i < e.PopulationSize; i++ {
		g := e.RandomGenome(random)
		population.Genomes[i] = g
	}
	population.SetFitness(e.Fitness)
	return population
}

func (s *Species[T]) RandomGenome(random Randomizer) *Genome[T] {
	var genome Genome[T]
	genome.Value = make([]T, s.GenomeSize)
	for i := 0; i < s.GenomeSize; i++ {
		genome.Value[i] = s.ValidGenes[random.Intn(len(s.ValidGenes))]
	}
	return &genome
}
func (e *Species[T]) CrossOver(random Randomizer, a, b *Genome[T]) *Genome[T] {
	dice := random.Float64()
	var child *Genome[T]
	switch {
	case dice < e.SingleCrossOverRate:
		child = e.SinglePointCrossOver(random, a, b)
	default:
		child = e.UniformCrossOver(random, a, b)
	}
	e.Mutate(random, child)
	return child
}

func (e *Species[T]) Mutate(random Randomizer, genome *Genome[T]) {
	for i := 0; i < e.GenomeSize; i++ {
		if random.Float64() < e.MutationRate {
			randomGene := random.Intn(len(e.ValidGenes))
			genome.Value[i] = e.ValidGenes[randomGene]
		}
	}
}

func (e *Species[T]) UniformCrossOver(random Randomizer, parentA, parentB *Genome[T]) *Genome[T] {
	next := Genome[T]{
		Value: make([]T, e.GenomeSize),
	}
	for i := 0; i < e.GenomeSize; i++ {
		switch random.Intn(2) {
		case 0:
			next.Value[i] = parentA.Value[i]
		case 1:
			next.Value[i] = parentB.Value[i]
		}
	}
	return &next
}

func (e *Species[T]) SinglePointCrossOver(random Randomizer, parentA, parentB *Genome[T]) *Genome[T] {
	crossOverPoint := random.Intn(e.GenomeSize)
	var next Genome[T]
	next.Value = make([]T, e.GenomeSize)
	for i := 0; i < e.GenomeSize; i++ {
		switch {
		case i > crossOverPoint:
			next.Value[i] = parentA.Value[i]
		default:
			next.Value[i] = parentB.Value[i]
		}
	}
	return &next
}

func (e *Species[T]) Selection(random Randomizer, population Population[T]) []*Genome[T] {
	var pool []*Genome[T]
	for i := 0; i < e.PopulationSize; i++ {
		num := int((population.Genomes[i].Fitness / population.Fittest.Fitness) * 100)
		for n := 0; n < num; n++ {
			pool = append(pool, population.Genomes[i])
		}
	}
	return pool
}

type Genome[T any] struct {
	Value   []T
	Fitness float64
}

func (genome *Genome[T]) SetFitness(f FitnessFunc[T]) {
	genome.Fitness = f(genome)
}

type Population[T any] struct {
	Genomes    []*Genome[T]
	Generation int
	Fittest    *Genome[T]
}

func (p *Population[T]) SetFitness(f FitnessFunc[T]) {
	for _, g := range p.Genomes {
		g.SetFitness(f)
		switch {
		case p.Fittest == nil:
			p.Fittest = g
		case p.Fittest.Fitness < g.Fitness:
			p.Fittest = g
		}
	}
}
