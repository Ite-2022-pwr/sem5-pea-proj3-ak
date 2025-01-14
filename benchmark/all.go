// benchmark jest odpowiedzialny za testowanie algorytmów
package benchmark

import (
	"fmt"
	"log"
	"path/filepath"
	"pea3/atsp"
	"pea3/generator"
	"pea3/utils"
	"runtime/debug"
)

// All - testuje oba algorytmy dla różnych ilości miast
func All() {
	promptBB := utils.BlueColor("Branch and Bound")
	promptDP := utils.BlueColor("Dynamic programming")

	totalTimeBB, totalTimeDP := 0.0, 0.0 // zmienne do przechowywania czasu rozwiązania
	var resultBB, resultDP [][]string    // wyniki

	log.Println(utils.BlueColor("[+] Rozpoczynanie testowania Algorytmów"))
	for numOfCities := MinVertices; numOfCities <= MaxVertices; numOfCities++ {
		for i := 0; i < NumberOfGraphs; i++ {
			log.Println(utils.BlueColor(fmt.Sprintf("Miast: %d, test: %d/%d", numOfCities, i+1, NumberOfGraphs)))
			G, _ := generator.GenerateAdjacencyMatrix(numOfCities)
			var tsp atsp.ATSP
			tsp = atsp.NewBranchAndBoundSolver(G)
			totalTimeBB += MeasureSolveTime(tsp, promptBB)

			tsp = atsp.NewDynamicProgrammingSolver(G)
			totalTimeDP += MeasureSolveTime(tsp, promptDP)
			debug.FreeOSMemory()
		}

		// średni czas dla każdego z algorytmów
		avgTimeBB := totalTimeBB / float64(NumberOfGraphs)
		avgTimeDP := totalTimeDP / float64(NumberOfGraphs)
		resultBB = append(resultBB, []string{fmt.Sprintf("%d", numOfCities), fmt.Sprintf("%d", int64(avgTimeBB))})
		resultDP = append(resultDP, []string{fmt.Sprintf("%d", numOfCities), fmt.Sprintf("%d", int64(avgTimeDP))})
		utils.SaveCSV(filepath.Join(OutputDirectory, "branch_and_bound3.csv"), resultBB)
		utils.SaveCSV(filepath.Join(OutputDirectory, "dynamic_programming3.csv"), resultDP)
		totalTimeBB, totalTimeDP = 0.0, 0.0
	}

	utils.SaveCSV(filepath.Join(OutputDirectory, "branch_and_bound3.csv"), resultBB)
	utils.SaveCSV(filepath.Join(OutputDirectory, "dynamic_programming3.csv"), resultDP)
}

func AllLocalSearch() {
	promptTS := utils.BlueColor("Tabu Search")
	promptSA := utils.BlueColor("Simulated Annealing")

	totalTimeTS, totalTimeSA := 0.0, 0.0 // zmienne do przechowywania czasu rozwiązania
	var resultTS, resultSA [][]string    // wyniki

	log.Println(utils.BlueColor("[+] Rozpoczynanie testowania Algorytmów"))
	for numOfCities := 50; numOfCities <= 250; numOfCities += 25 {
		for i := 0; i < NumberOfGraphs; i++ {
			log.Println(utils.BlueColor(fmt.Sprintf("Miast: %d, test: %d/%d", numOfCities, i+1, NumberOfGraphs)))
			G, _ := generator.GenerateAdjacencyMatrix(numOfCities)
			var tsp atsp.ATSP
			tsp = atsp.NewTabuSearchSolver(G, 20, 1000, atsp.MovingInsert)

			totalTimeTS += MeasureSolveTime(tsp, promptTS)

			tsp = atsp.NewSimulatedAnnealingSolver(G, 0.975, 1e-12, 9000, 4000)
			totalTimeSA += MeasureSolveTime(tsp, promptSA)
			debug.FreeOSMemory()
		}

		// średni czas dla każdego z algorytmów
		avgTimeTS := totalTimeTS / float64(NumberOfGraphs)
		avgTimeSA := totalTimeSA / float64(NumberOfGraphs)
		resultTS = append(resultTS, []string{fmt.Sprintf("%d", numOfCities), fmt.Sprintf("%.3f", avgTimeTS/1000000000.0)})
		resultSA = append(resultSA, []string{fmt.Sprintf("%d", numOfCities), fmt.Sprintf("%.3f", avgTimeSA/1000000000.0)})
		utils.SaveCSV(filepath.Join(OutputDirectory, "tabu_search2.csv"), resultTS)
		utils.SaveCSV(filepath.Join(OutputDirectory, "simulated_annealing2.csv"), resultSA)
		totalTimeTS, totalTimeSA = 0.0, 0.0
	}

	utils.SaveCSV(filepath.Join(OutputDirectory, "tabu_search2.csv"), resultTS)
	utils.SaveCSV(filepath.Join(OutputDirectory, "simulated_annealing2.csv"), resultSA)
}

func AllMetaheuristicTime() {
	var result [][]string

	log.Println(utils.BlueColor("[+] Rozpoczynanie testowania Dynamic Programming"))
	for numOfCities := MinVertices; numOfCities <= MaxVertices; numOfCities++ {
		var tsp atsp.ATSP
		totalTimeTS, totalTimeSA, totalTimeACO := 0.0, 0.0, 0.0

		for i := 0; i < Rounds; i++ {
			log.Println(utils.BlueColor(fmt.Sprintf("Miast: %d, %d/%d", numOfCities, i+1, Rounds)))
			G, _ := generator.GenerateAdjacencyMatrix(numOfCities)

			tsp = atsp.NewTabuSearchSolver(G, 20, 1000, atsp.MovingInsert)
			totalTimeTS += MeasureSolveTime(tsp, "Tabu Search")

			tsp = atsp.NewSimulatedAnnealingSolver(G, 0.975, 1e-12, 9000, 4000)
			totalTimeSA += MeasureSolveTime(tsp, "Simulated Annealing")

			tsp = atsp.NewAntColonyOptimizationSolver(G, 10, 1, 3, 20, 0.25, 50)
			totalTimeACO += MeasureSolveTime(tsp, "Ant Colony Optimization")
		}
		avgTimeTS := totalTimeTS / float64(Rounds)
		avgTimeSA := totalTimeSA / float64(Rounds)
		avgTimeACO := totalTimeACO / float64(Rounds)

		result = append(result, []string{
			fmt.Sprintf("%d", numOfCities),
			fmt.Sprintf("%f", avgTimeTS/1000000000.0),
			fmt.Sprintf("%f", avgTimeSA/1000000000.0),
			fmt.Sprintf("%f", avgTimeACO/1000000000.0),
		})
		utils.SaveCSV(filepath.Join(OutputDirectory, "COMPARE_TIME.csv"), result)
		debug.FreeOSMemory()
	}

	utils.SaveCSV(filepath.Join(OutputDirectory, "COMPARE_TIME.csv"), result)
}

func AllMetaheuristicError() {
	var result [][]string

	log.Println(utils.BlueColor("[+] Rozpoczynanie testowania Dynamic Programming"))
	for numOfCities := MinVertices; numOfCities <= MaxVertices; numOfCities++ {
		var tsp atsp.ATSP
		log.Println(utils.BlueColor(fmt.Sprintf("Miast: %d", numOfCities)))
		G, _ := generator.GenerateAdjacencyMatrix(numOfCities)
		tsp = atsp.NewDynamicProgrammingSolver(G)
		_, bestCost := MeasureSolveTimeWithCost(tsp, "Dynamic Programming")
		log.Println(utils.BlueColor(fmt.Sprintf("Optymalny wynik: %d", bestCost)))
		debug.FreeOSMemory()

		tsp = atsp.NewTabuSearchSolver(G, 20, 1000, atsp.MovingInsert)
		timeTS, costTS := MeasureSolveTimeWithCost(tsp, "Tabu Search")

		tsp = atsp.NewSimulatedAnnealingSolver(G, 0.975, 1e-12, 9000, 4000)
		timeSA, costSA := MeasureSolveTimeWithCost(tsp, "Simulated Annealing")

		tsp = atsp.NewAntColonyOptimizationSolver(G, 10, 1, 3, 20, 0.25, 50)
		timeACO, costACO := MeasureSolveTimeWithCost(tsp, "Ant Colony Optimization")

		result = append(result, []string{
			fmt.Sprintf("%d", numOfCities),
			fmt.Sprintf("%d", bestCost),
			fmt.Sprintf("%d", costTS),
			fmt.Sprintf("%d", CalculateError(costTS, bestCost)),
			fmt.Sprintf("%f", timeTS),
			fmt.Sprintf("%d", costSA),
			fmt.Sprintf("%d", CalculateError(costSA, bestCost)),
			fmt.Sprintf("%f", timeSA),
			fmt.Sprintf("%d", costACO),
			fmt.Sprintf("%d", CalculateError(costACO, bestCost)),
			fmt.Sprintf("%f", timeACO),
		})
		utils.SaveCSV(filepath.Join(OutputDirectory, "COMPARE.csv"), result)
		debug.FreeOSMemory()
	}

	utils.SaveCSV(filepath.Join(OutputDirectory, "COMPARE.csv"), result)
}
