package teamform

import (
	"fmt"
	"math"
	"math/rand"
)

// SimulatedAnnealingResult is the result of the SimulatedAnnealing function
type SimulatedAnnealingResult struct {
	Groups                                  Teams
	CandidateScoreHistory, BestScoreHistory []float64
}

// SA is the Simulated Annealing algorithm for grouping the students into teams
// random is the random number generator
// iterations is the number of iterations to perform the algorithm - the more the longer it takes, but the better the result
// stepPermutations is the number of permutations to perform on each iteration to generate a candidate - should be 1
// temperature is the initial temperature of the algorithm. The higher the temperature, the more likely it is to accept a worse candidate at the beginning
// debug is a boolean which indicates whether to print debug information
// wSatisfaction, wC, wE, wN are the weights for the objective function
func SA(groups Teams,
	random *rand.Rand,
	iterations int,
	stepPermutations int,
	temperature float64,
	debug bool,
	wSatisfaction, wC, wE, wN float64,
) SimulatedAnnealingResult {

	current, currentEval := groups, Objective(groups, wSatisfaction, wC, wE, wN)
	best, bestEval := current, currentEval

	var candidateEval, currentTemperature, metropolis, diff, randomValue float64

	candidateScoreHistory := make([]float64, iterations)
	bestScoreHistory := make([]float64, iterations)

	// Pre-allocate the reverse movements slice
	reverseMovements := NewMovementsSlice(stepPermutations)

	for i := 0; i < iterations; i++ {

		// Perform a permutation on the groups and get the reverse movements to be able to revert the permutation
		// The 'Current' variable now functions as a 'Candidate' variable
		current.ShuffleStudents(random, reverseMovements)

		// Calculate the objective function for the candidate
		candidateEval = Objective(current, wSatisfaction, wC, wE, wN)

		if candidateEval > bestEval {
			// Thee onlu way to acually preserve the best is to copy the slice
			// We could keep a track of the movments, which might be more efficient
			// , but could perhaps be more error prone
			bestCopy := make([]Team, len(current))

			for i, team := range current {
				bestCopy[i] = Team{
					Students: make([]Student, len(team.Students)),
					S:        team.S,
					C:        team.C,
					E:        team.E,
					N:        team.N,
				}
				copy(bestCopy[i].Students, team.Students)
			}

			best = bestCopy
			bestEval = candidateEval
		}

		// Calculate the difference between the candidate and the current
		diff = candidateEval - currentEval

		// Calculate the current temperature
		currentTemperature = temperature / float64(i+1)

		// Calculate the metropolis value
		metropolis = math.Exp(diff / currentTemperature)

		// Get a random value between 0 and 1
		randomValue = random.Float64()

		// If the difference is positive or the random value is smaller than the metropolis value, accept the candidate
		if diff > 0 || randomValue < metropolis {
			if debug {
				if diff > 0 {
					fmt.Printf("🟢 [%d] diff = %.2f metro = %.2f, temp = %.2f \n", i, diff, metropolis, currentTemperature)
				} else {
					fmt.Printf("🔴 [%d] metro = exp ((%.2f - %.2f) / %.2f) = %.2f, which is more than rnd = %.2f\n",
						i,
						candidateEval,
						currentEval,
						currentTemperature,
						metropolis,
						randomValue,
					)
				}
			}

			// We implicitly set the candidate (pointers): we only update the evaluation
			currentEval = candidateEval
		} else {
			if debug {
				fmt.Printf("⚪️ [%d] Didn't apply any changes, reverting the list\n", i)
			}

			// With the reverse movements, we can revert the permutation/ shuffles
			// With this we reverse the 'Candidate' back to the 'Current'
			current.ReverseStudentShuffles(reverseMovements)
		}

		candidateScoreHistory[i] = currentEval
		bestScoreHistory[i] = bestEval
	}

	return SimulatedAnnealingResult{
		Groups:                best,
		CandidateScoreHistory: candidateScoreHistory,
		BestScoreHistory:      bestScoreHistory,
	}
}
