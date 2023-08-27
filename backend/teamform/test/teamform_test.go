package teamform_test

import (
	"fmt"
	"image/color"
	"jv/team-tone-tuner/teamform"
	"math/rand"
	"testing"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// ------------------- Tests -------------------

// Test that students are correctly shuffled between teams
// and that the reverse movements are correct
func TestTeamPermutations(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	teams := getDefaultTestingTeams(r)

	originalGroups := teams.ToString()

	println("Original groups: ")
	println(originalGroups)

	reverseMovements := teamform.NewMovementsSlice(10000)

	teams.ShuffleStudents(r, reverseMovements)

	permutedGroups := teams.ToString()

	println("Groups after permutation: ")
	println(permutedGroups)

	teams.ReverseStudentShuffles(reverseMovements)

	reversedGroups := teams.ToString()

	println("Should be the same: ")
	println(reversedGroups)

	if permutedGroups == originalGroups {
		t.Errorf("Groups were not permuted.\nOriginal: %s\nPermuted: %s", originalGroups, permutedGroups)
	}

	if originalGroups != reversedGroups {
		t.Errorf("Groups were not reversed back to original state.\nOriginal: %s\nReversed: %s", originalGroups, reversedGroups)
	}
}

// TextXGBoost tests the XGBoost model - if it loads correctly and if it can predict the satisfaction
func TestXGBoost(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	teams := getDefaultTestingTeams(r)

	t.Log(teamform.SatisfactionObjective(teams))
}

func TestObjective(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	teams := getDefaultTestingTeams(r)

	objective := teamform.Objective(teams, 0.5, 2, 1, 1)

	if objective == 0 {
		t.Errorf("Objective function returned 0")
	}

	fmt.Printf("Objective function returned %f \n", objective)
}

func TestAll(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	groups := generateRandomTeams(20, 5, r)

	fmt.Printf("Before obj %f \n", teamform.Objective(groups, 0.5, 2, 1, 1))
	fmt.Println(groups.ToString())

	timer := time.Now()

	res := teamform.SA(groups, r, 10000, 1, 200.0, false, 0.5, 2, 1, 1)

	fmt.Printf("Time: %s \n", time.Since(timer))

	p := plot.New()

	p.Title.Text = "Test"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	lineDataBestScore := transformToPoints(res.BestScoreHistory)
	lineDataCandidateScore := transformToPoints(res.CandidateScoreHistory)

	l1, err := plotter.NewLine(lineDataBestScore)
	l1.Color = color.RGBA{R: 255, A: 255}

	if err != nil {
		panic(err)
	}

	l2, err := plotter.NewLine(lineDataCandidateScore)
	l2.Color = color.RGBA{B: 255, A: 120}

	if err != nil {
		panic(err)
	}

	p.Add(l1, l2)

	// Save the plot to a PNG file.
	if err := p.Save(20*vg.Inch, 8*vg.Inch, "points.png"); err != nil {
		panic(err)
	}

	fmt.Printf("After obj %f \n", teamform.Objective(groups, 0.5, 2, 1, 1))
	fmt.Println(groups.ToString())
}

// transformToPoints takes a slice of float64 and transforms it into a slice of plotter.XYs
func transformToPoints(series []float64) plotter.XYs {
	pts := make(plotter.XYs, len(series))

	for i, v := range series {
		pts[i].X = float64(i)
		pts[i].Y = v
	}

	return pts
}
