package teamform

import (
	"fmt"
	"jv/team-tone-tuner/model"
	"math/rand"
)

func randomPairWithoutDuplicates(random *rand.Rand, start int, end int) (int, int) {
	if end <= start {
		panic(fmt.Sprintf("end (%d) must be greater than start (%d)", end, start))
	}

	num1 := random.Intn(end-start) + start
	num2 := random.Intn(end-start-1) + start

	if num2 >= num1 {
		num2++
	}

	return num1, num2
}

// convertUserToStudent converts a user to a student used for the SA algorithm
func convertUserToStudent(user *model.User) (Student, error) {

	if user == nil {
		return Student{}, fmt.Errorf("User cannot be nil")
	}

	if user.Edges.BfiReport == nil {
		return Student{}, fmt.Errorf("User must have a BFI report")
	}

	rep := user.Edges.BfiReport

	s := Student{
		GithubNick: user.GithubUsername,
		Id:         user.ID,
		O:          rep.Openness.PointsNormalized,
		C:          rep.Conscientiousness.PointsNormalized,
		E:          rep.Extraversion.PointsNormalized,
		A:          rep.Agreeableness.PointsNormalized,
		N:          rep.Neuroticism.PointsNormalized,
	}

	return s, nil
}
