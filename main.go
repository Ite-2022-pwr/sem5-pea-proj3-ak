package main

import (
	"log"
	"pea3/atsp"
	"pea3/benchmark"
	"pea3/utils"
)

func main() {
	//menu.RunMenu()

	G, err := utils.ReadGraphFromFile("data/input/ftv170.txt")

	if err != nil {
		log.Fatal(err)
	}

	tsp := atsp.NewAntColonyOptimizationSolver(G, 3, 1, 3, 100, 0.3, 1.0, 100.0)

	benchmark.MeasureSolveTime(tsp, "ACO")
}
