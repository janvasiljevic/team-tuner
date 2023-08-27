package teamform

import (
	"bufio"
	"bytes"
	"fmt"
	"jv/team-tone-tuner/xgb"

	_ "embed"

	"github.com/dmitryikh/leaves"
)

var (
	xgbModel *leaves.Ensemble
)

// Create an init function to load the model
func init() {
	var err error

	reader := bytes.NewReader(xgb.ModelBytes)

	print(reader.Len())

	buf := bufio.NewReader(reader) // Convert bytes.Reader to bufio.Reader

	xgbModel, err = leaves.XGEnsembleFromReader(buf, true)

	if err != nil {
		panic(err)
	}
}

const (
	// Assumed amount of work is the baseline to predict the satisfaction
	ASSUMED_AOW = 0.80
)

func PrintModelInfo() {
	fmt.Printf("Name: %s\n", xgbModel.Name())
	fmt.Printf("NFeatures: %d\n", xgbModel.NFeatures())
	fmt.Printf("NOutputGroups: %d\n", xgbModel.NOutputGroups())
	fmt.Printf("NEstimators: %d\n", xgbModel.NEstimators())
	fmt.Printf("Transformation: %s\n", xgbModel.Transformation().Name())
	fmt.Println(xgbModel.NEstimators())
}

func SatisfactionObjective(teams Teams) float64 {

	// Where we store data to input to xgboost
	denseMatrix := make([]float64, 0)

	// Team mapping matrix, so we can re assemble the outputs
	// of xgboost and calculate per team satisfaction
	// needed for standard deviation
	// Should be 11 times smaller than denseMatrix
	studentTeamMapping := make([]int, 0)

	recalculateTeams := make([]*Team, 0)

	for teamIdx := range teams {
		// No need to calculate the satisfaction for teams that are not dirty

		if !teams[teamIdx].Dirty {
			continue
		}

		// Don't know how much time I spent on this bug here
		// PLEASE REMEMBER, THAT GO MAKES A COPY WITH RANGE
		recalculateTeams = append(recalculateTeams, &teams[teamIdx])

		for s1Idx, s1 := range teams[teamIdx].Students {
			for s2Idx, s2 := range teams[teamIdx].Students {
				// do not compare self to self
				if s1Idx == s2Idx {
					continue
				}

				// fmt.Printf("%s - %s\n", s1.GithubNick, s2.GithubNick)
				s1.AddOCEANToDenseMatrix(&denseMatrix)
				s2.AddOCEANToDenseMatrix(&denseMatrix)
				denseMatrix = append(denseMatrix, ASSUMED_AOW)
				studentTeamMapping = append(studentTeamMapping, teamIdx)
			}
		}
	}

	predictions := make([]float64, len(studentTeamMapping))

	err := xgbModel.PredictDense(
		denseMatrix,             // data presented in a dense matrix
		len(studentTeamMapping), // number of vectors (rows)
		11,                      // number of features (columns)
		predictions,             // where we saving the results
		130,                     // estimators
		4,                       // num of cores
	)

	if err != nil {
		panic(err)
	}

	// Reset the teams (dirty flag = false, and scores to 0)
	for _, team := range recalculateTeams {
		team.Reset()
	}

	// loop over the team predictions and add coresponding scores to the team
	// keep in mind we have a binary classification model, so we need to
	// convert the scores to 0 or 1
	for i, teamIndex := range studentTeamMapping {
		score := predictions[i]

		if score > 0.5 {
			teams[teamIndex].S += 1
		}
	}

	avg := 0.0
	std := 0.0

	for _, team := range teams {
		avg += team.S
	}

	avg /= float64(len(teams))

	for _, team := range teams {
		std += (team.S - avg) * (team.S - avg)
	}

	std /= float64(len(teams))

	return avg - std*2
}

// ObjectivePerDimensions calculates the objective function per dimension
// Calculates score for a given dimension (e.g. conscientiousness) for a given set of teams.
// The score is calculated as the average of the standard deviations of the dimension for each team.
// It tries to maximize the score, while keeping the teams balanced.
// If `minimizeScore` is True, it tries to minimize the score, while still keeping the teams balanced.
func ObjectivePerDimensions(teams Teams, getter func(Student) float64, minimizeScore bool) float64 {
	avgs := make([]float64, len(teams))

	for i, team := range teams {
		for _, student := range team.Students {
			avgs[i] += getter(student)
		}

		avgs[i] /= float64(len(team.Students))
	}

	avg := 0.0

	for _, a := range avgs {
		avg += a
	}

	avg /= float64(len(avgs))

	std := 0.0

	for _, a := range avgs {
		std += (a - avg) * (a - avg)
	}

	std /= float64(len(avgs))

	if minimizeScore {
		return 1 - avg + std
	}

	return avg - std
}

func Objective(teams Teams, wSatisfaction, wC, wE, wN float64) float64 {

	c := ObjectivePerDimensions(teams, func(s Student) float64 { return s.C }, true)
	e := ObjectivePerDimensions(teams, func(s Student) float64 { return s.E }, true)
	n := ObjectivePerDimensions(teams, func(s Student) float64 { return s.N }, false)
	sat := SatisfactionObjective(teams)

	return wSatisfaction*sat + wC*c + wE*e + wN*n
}
