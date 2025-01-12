package atsp

import (
	"math"
	"math/rand"
	"pea3/graph"
	"runtime/debug"
)

type AntColonyOptimizationSolver struct {
	graph graph.Graph

	ants                 int     // liczba mrówek
	alpha                int     // waga feromonów
	beta                 int     // waga heurystyki
	iterations           int     // liczba iteracji
	pheromoneEvaporation float64 // współczynnik odparowania feromonów
	pheromoneInit        float64 // początkowa wartość feromonów
	q                    float64 // stała do obliczania ilości pozostawionych feromonów
}

func NewAntColonyOptimizationSolver(G graph.Graph, ants, alpha, beta, iterations int, pheromoneEvaporation, pheromoneInit, q float64) *AntColonyOptimizationSolver {
	return &AntColonyOptimizationSolver{
		graph:                G,
		ants:                 ants,
		alpha:                alpha,
		beta:                 beta,
		iterations:           iterations,
		pheromoneEvaporation: pheromoneEvaporation,
		pheromoneInit:        pheromoneInit,
		q:                    q,
	}
}

func (aco *AntColonyOptimizationSolver) Solve(startVertex int) (int, []int) {
	return aco.AntColonyOptimization(startVertex)
}

func (aco *AntColonyOptimizationSolver) GetGraph() graph.Graph {
	return aco.graph
}

func (aco *AntColonyOptimizationSolver) AntColonyOptimization(startVertex int) (int, []int) {
	bestCost, bestPath := math.MaxInt, []int{}

	pheromones := aco.initializePheromones()

	for iter := 0; iter < aco.iterations; iter++ {
		allPaths := make([][]int, aco.ants)
		allCosts := make([]int, aco.ants)

		for ant := 0; ant < aco.ants; ant++ {
			cost, path := aco.findPath(pheromones)
			debug.FreeOSMemory()
			allPaths[ant] = path
			allCosts[ant] = cost
			if cost < bestCost {
				bestCost, bestPath = cost, path
			}
		}

		aco.updatePheromones(pheromones, allPaths, allCosts)
	}

	return bestCost, bestPath
}

func (aco *AntColonyOptimizationSolver) initializePheromones() [][]float64 {
	pheromones := make([][]float64, aco.GetGraph().GetVerticesCount())
	for i := 0; i < aco.GetGraph().GetVerticesCount(); i++ {
		pheromones[i] = make([]float64, aco.GetGraph().GetVerticesCount())
		for j := 0; j < aco.GetGraph().GetVerticesCount(); j++ {
			pheromones[i][j] = aco.pheromoneInit
		}
	}
	return pheromones
}

func (aco *AntColonyOptimizationSolver) findPath(pheromones [][]float64) (int, []int) {
	path := make([]int, aco.GetGraph().GetVerticesCount())
	visited := make([]bool, aco.GetGraph().GetVerticesCount())
	path[0] = rand.Intn(aco.GetGraph().GetVerticesCount())
	visited[path[0]] = true
	pathCost := 0

	for i := 0; i < aco.GetGraph().GetVerticesCount()-1; i++ {
		nextVertex := aco.selectNextVertex(path[i], visited, pheromones)
		debug.FreeOSMemory()
		path[i+1] = nextVertex
		pathCost += aco.GetGraph().AsMatrix()[path[i]][nextVertex]
		visited[nextVertex] = true
	}
	pathCost += aco.GetGraph().AsMatrix()[path[aco.GetGraph().GetVerticesCount()-1]][path[0]]
	return pathCost, path
}

// selectNextVertex znajduje następne miasto do odwiedzenia metodą weighted random algorithm: https://dev.to/jacktt/understanding-the-weighted-random-algorithm-581p
func (aco *AntColonyOptimizationSolver) selectNextVertex(currentVertex int, visited []bool, pheromones [][]float64) int {
	probabilities := make([]float64, aco.GetGraph().GetVerticesCount())
	sum := 0.0

	for i := 0; i < aco.GetGraph().GetVerticesCount(); i++ {
		if !visited[i] {
			probabilities[i] = math.Pow(pheromones[currentVertex][i], float64(aco.alpha)) * math.Pow(1.0/float64(aco.GetGraph().AsMatrix()[currentVertex][i]), float64(aco.beta))
			sum += probabilities[i]
		}
	}

	randomVal := rand.Float64() * sum
	cumulative := 0.0

	for i := 0; i < aco.GetGraph().GetVerticesCount(); i++ {
		if !visited[i] {
			cumulative += probabilities[i]
			if cumulative >= randomVal {
				return i
			}
		}
	}

	return -1
}

func (aco *AntColonyOptimizationSolver) updatePheromones(pheromones [][]float64, allPaths [][]int, allCosts []int) {
	for i := 0; i < aco.GetGraph().GetVerticesCount(); i++ {
		for j := 0; j < aco.GetGraph().GetVerticesCount(); j++ {
			pheromones[i][j] *= (1 - aco.pheromoneEvaporation)
		}
	}

	for i := 0; i < aco.ants; i++ {
		for j := 0; j < aco.GetGraph().GetVerticesCount()-1; j++ {
			pheromones[allPaths[i][j]][allPaths[i][j+1]] += aco.q / float64(allCosts[i])
		}
		pheromones[allPaths[i][aco.GetGraph().GetVerticesCount()-1]][allPaths[i][0]] += aco.q / float64(allCosts[i])
	}
}
