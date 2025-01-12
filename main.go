package main

import (
	"log"
	"pea3/atsp"
	"pea3/benchmark"
	"pea3/utils"
)

func main() {
	//menu.RunMenu()

	G, err := utils.ReadGraphFromFile("data/input/ftv55.txt")

	if err != nil {
		log.Fatal(err)
	}

	tsp := atsp.NewAntColonyOptimizationSolver(G, G.GetVerticesCount(), 1, 4, 20, 0.5, 100.0)

	benchmark.MeasureSolveTime(tsp, "ACO")
}
