package benchmark

import (
	"fmt"
	"log"
	"path/filepath"
	"pea3/atsp"
	"pea3/utils"
)

var (
	ants          = []int{10, 20, 30}
	alphas        = []int{1, 2, 3}
	betas         = []int{2, 3, 4}
	evaporations  = []float64{0.25, 0.5, 0.75}
	acoIterations = []int{20, 60, 100}
	qs            = []float64{50, 75, 100}
)

func TestACOAnts() {
	prompt := "Ant Colony Optimization"

	for filename, optimalCost := range OptimalSolutionsACO {
		G, err := utils.ReadGraphFromFile(filepath.Join(InputDirectory, filename))
		if err != nil {
			log.Fatal(utils.RedColor(err))
		}

		// Testowanie liczby mrówek
		for a := 0; a < len(ants); a++ {
			var result [][]string
			for r := 0; r < Rounds; r++ {
				log.Println(utils.BlueColor(fmt.Sprintf("ACO ants: %v, test: %d/%d", ants[a], r+1, Rounds)))
				tsp := atsp.NewAntColonyOptimizationSolver(G, ants[a], alphas[0], betas[0], acoIterations[0], evaporations[0], qs[0])
				elapsed, cost := MeasureSolveTimeWithCost(tsp, prompt)
				result = append(result, []string{
					"ANTS",
					fmt.Sprintf("%d", ants[a]),
					fmt.Sprintf("%d", cost),
					fmt.Sprintf("%d", CalculateError(cost, optimalCost)),
					fmt.Sprintf("%.3f", elapsed/1000000000.0),
				})
				utils.SaveCSV(filepath.Join(OutputDirectory, fmt.Sprintf("ACO_ANTS_%v_%d.csv", filename, ants[a])), result)
			}
		}
	}
}

func TestACOAlphas() {
	prompt := "Ant Colony Optimization"

	for filename, optimalCost := range OptimalSolutionsACO {
		G, err := utils.ReadGraphFromFile(filepath.Join(InputDirectory, filename))
		if err != nil {
			log.Fatal(utils.RedColor(err))
		}

		// Testowanie alfy
		for a := 0; a < len(alphas); a++ {
			var result [][]string
			for r := 0; r < Rounds; r++ {
				log.Println(utils.BlueColor(fmt.Sprintf("ACO alpha: %v, test: %d/%d", alphas[a], r+1, Rounds)))
				tsp := atsp.NewAntColonyOptimizationSolver(G, ants[0], alphas[a], betas[0], acoIterations[0], evaporations[0], qs[0])
				elapsed, cost := MeasureSolveTimeWithCost(tsp, prompt)
				result = append(result, []string{
					"ALPHA",
					fmt.Sprintf("%d", alphas[a]),
					fmt.Sprintf("%d", cost),
					fmt.Sprintf("%d", CalculateError(cost, optimalCost)),
					fmt.Sprintf("%.3f", elapsed/1000000000.0),
				})
				utils.SaveCSV(filepath.Join(OutputDirectory, fmt.Sprintf("ACO_ALPHA_%v_%d.csv", filename, alphas[a])), result)
			}
		}
	}
}

func TestACOBetas() {
	prompt := "Ant Colony Optimization"

	for filename, optimalCost := range OptimalSolutionsACO {
		G, err := utils.ReadGraphFromFile(filepath.Join(InputDirectory, filename))
		if err != nil {
			log.Fatal(utils.RedColor(err))
		}

		// Testowanie bety
		for b := 0; b < len(alphas); b++ {
			var result [][]string
			for r := 0; r < Rounds; r++ {
				log.Println(utils.BlueColor(fmt.Sprintf("ACO beta: %v, test: %d/%d", betas[b], r+1, Rounds)))
				tsp := atsp.NewAntColonyOptimizationSolver(G, ants[0], alphas[0], betas[b], acoIterations[0], evaporations[0], qs[0])
				elapsed, cost := MeasureSolveTimeWithCost(tsp, prompt)
				result = append(result, []string{
					"BETA",
					fmt.Sprintf("%d", betas[b]),
					fmt.Sprintf("%d", cost),
					fmt.Sprintf("%d", CalculateError(cost, optimalCost)),
					fmt.Sprintf("%.3f", elapsed/1000000000.0),
				})
				utils.SaveCSV(filepath.Join(OutputDirectory, fmt.Sprintf("ACO_BETA_%v_%d.csv", filename, betas[b])), result)
			}
		}
	}
}

func TestACOEvaporations() {
	prompt := "Ant Colony Optimization"

	for filename, optimalCost := range OptimalSolutionsACO {
		G, err := utils.ReadGraphFromFile(filepath.Join(InputDirectory, filename))
		if err != nil {
			log.Fatal(utils.RedColor(err))
		}

		// Testowanie odparowywania feromonów
		for e := 0; e < len(evaporations); e++ {
			var result [][]string
			for r := 0; r < Rounds; r++ {
				log.Println(utils.BlueColor(fmt.Sprintf("ACO evaporation: %v, test: %d/%d", evaporations[e], r+1, Rounds)))
				tsp := atsp.NewAntColonyOptimizationSolver(G, ants[0], alphas[0], betas[0], acoIterations[0], evaporations[e], qs[0])
				elapsed, cost := MeasureSolveTimeWithCost(tsp, prompt)
				result = append(result, []string{
					"EVAPORATION",
					fmt.Sprintf("%f", evaporations[e]),
					fmt.Sprintf("%d", cost),
					fmt.Sprintf("%d", CalculateError(cost, optimalCost)),
					fmt.Sprintf("%.3f", elapsed/1000000000.0),
				})
				utils.SaveCSV(filepath.Join(OutputDirectory, fmt.Sprintf("ACO_EVAPORATION_%v_%f.csv", filename, evaporations[e])), result)
			}
		}
	}
}

func TestACOIterations() {
	prompt := "Ant Colony Optimization"

	for filename, optimalCost := range OptimalSolutionsACO {
		G, err := utils.ReadGraphFromFile(filepath.Join(InputDirectory, filename))
		if err != nil {
			log.Fatal(utils.RedColor(err))
		}

		// Testowanie liczby iteracji
		for i := 0; i < len(acoIterations); i++ {
			var result [][]string
			for r := 0; r < Rounds; r++ {
				log.Println(utils.BlueColor(fmt.Sprintf("ACO iterations: %v, test: %d/%d", acoIterations[i], r+1, Rounds)))
				tsp := atsp.NewAntColonyOptimizationSolver(G, ants[0], alphas[0], betas[0], acoIterations[i], evaporations[0], qs[0])
				elapsed, cost := MeasureSolveTimeWithCost(tsp, prompt)
				result = append(result, []string{
					"EVAPORATION",
					fmt.Sprintf("%d", acoIterations[i]),
					fmt.Sprintf("%d", cost),
					fmt.Sprintf("%d", CalculateError(cost, optimalCost)),
					fmt.Sprintf("%.3f", elapsed/1000000000.0),
				})
				utils.SaveCSV(filepath.Join(OutputDirectory, fmt.Sprintf("ACO_ITERATIONS_%v_%d.csv", filename, acoIterations[i])), result)
			}
		}
	}
}

func TestACOQs() {
	prompt := "Ant Colony Optimization"

	for filename, optimalCost := range OptimalSolutionsACO {
		G, err := utils.ReadGraphFromFile(filepath.Join(InputDirectory, filename))
		if err != nil {
			log.Fatal(utils.RedColor(err))
		}

		// Testowanie zostawiania feromonów przez mrówki
		for q := 0; q < len(qs); q++ {
			var result [][]string
			for r := 0; r < Rounds; r++ {
				log.Println(utils.BlueColor(fmt.Sprintf("ACO Q: %v, test: %d/%d", qs[q], r+1, Rounds)))
				tsp := atsp.NewAntColonyOptimizationSolver(G, ants[0], alphas[0], betas[0], acoIterations[0], evaporations[0], qs[q])
				elapsed, cost := MeasureSolveTimeWithCost(tsp, prompt)
				result = append(result, []string{
					"Q",
					fmt.Sprintf("%d", int(qs[q])),
					fmt.Sprintf("%d", cost),
					fmt.Sprintf("%d", CalculateError(cost, optimalCost)),
					fmt.Sprintf("%.3f", elapsed/1000000000.0),
				})
				utils.SaveCSV(filepath.Join(OutputDirectory, fmt.Sprintf("ACO_Q_%v_%d.csv", filename, int(qs[q]))), result)
			}
		}
	}
}
