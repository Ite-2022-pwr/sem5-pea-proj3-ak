package atsp

import "pea3/graph"

type AntColonyOptimizationSolver struct {
	graph graph.Graph

	ants           int     // liczba mrówek
	alpha          int     // waga feromonów
	beta           int     // waga odległości
	iterations     int     // liczba iteracji
	pheromoneDecay float64 // współczynnik odparowania feromonów
	pheromoneInit  float64 // początkowa wartość feromonów
	q              float64 // stała do obliczania ilości pozostawionych feromonów
}

func NewAntColonyOptimizationSolver(G graph.Graph, ants, alpha, beta, iterations int, pheromoneDecay, pheromoneInit, q float64) *AntColonyOptimizationSolver {
	return &AntColonyOptimizationSolver{
		graph:          G,
		ants:           ants,
		alpha:          alpha,
		beta:           beta,
		iterations:     iterations,
		pheromoneDecay: pheromoneDecay,
		pheromoneInit:  pheromoneInit,
		q:              q,
	}
}

func (aco *AntColonyOptimizationSolver) Solve(startVertex int) (int, []int) {
	return aco.AntColonyOptimization(startVertex)
}

func (aco *AntColonyOptimizationSolver) GetGraph() graph.Graph {
	return aco.graph
}

func (aco *AntColonyOptimizationSolver) AntColonyOptimization(startVertex int) (int, []int) {
	return -1, nil // not implemented
}
