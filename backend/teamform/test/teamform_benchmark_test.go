package teamform_test

import (
	"fmt"
	"jv/team-tone-tuner/teamform"
	"math/rand"
	"testing"
	"time"
)

// BenchmarkSimulateAnnealing is a benchmarking function for the SimulatedAnnealing function.
// It tests the function with different numbers of teams and iterations:
// The iterations and num of teams were chosen to be the same as the ones used in the
// Python implementation of the algorithm - to be compared in the thesis!.
func BenchmarkSimulateAnnealing(b *testing.B) {

	type BenchmarkStruct struct {
		iterations int
		numOfTeams int
	}

	benchmarkTable := []BenchmarkStruct{
		{numOfTeams: 10, iterations: 6000},
		{numOfTeams: 12, iterations: 6000},
		{numOfTeams: 14, iterations: 8000},
		{numOfTeams: 16, iterations: 10000},
		{numOfTeams: 18, iterations: 12000},
		{numOfTeams: 20, iterations: 12000},
		{numOfTeams: 22, iterations: 12000},
		{numOfTeams: 24, iterations: 16000},
		{numOfTeams: 26, iterations: 16000},
		{numOfTeams: 28, iterations: 20000},
		{numOfTeams: 30, iterations: 20000},
		{numOfTeams: 32, iterations: 20000},
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for _, item := range benchmarkTable {
		b.Run(fmt.Sprintf("teams: %d, iter: %d", item.numOfTeams, item.iterations), func(b_sub *testing.B) {
			groups := generateRandomTeams(item.numOfTeams, 5, r)

			teamform.SA(groups, r, item.iterations, 1, 12, false, 0.5, 2, 1, 1)
		})
	}
}

/*
These are the results of the benchmarking:

Teams,Iterations,Average Time (seconds)
	10,6000,1.4845
	12,6000,1.4610
	14,8000,1.9513
	16,10000,2.4625
	18,12000,2.9743
	20,12000,2.9978
	22,12000,2.9252
	24,16000,3.9164
	26,16000,3.9057
	28,20000,4.9783
	30,20000,4.9217
	32,20000,4.9235
*/
