package teamform

import (
	"fmt"
	"jv/team-tone-tuner/model"
	"math/rand"
	"sort"
	"time"
)

func GenerateTeams(users []*model.User,
	teamSizes []int,
	saIterations int,
	saTemperature float64,
	saWeightSatisfaction, saWeightN, saWeightE, saWeightC float64,
) (*SimulatedAnnealingResult, error) {
	numOfStudentsRequest := 0

	for _, groupSize := range teamSizes {
		numOfStudentsRequest += groupSize
	}

	if numOfStudentsRequest != len(users) {
		return nil, fmt.Errorf("Number of students in request does not match number of students in database")
	}

	convertedUsers := make([]Student, len(users))

	// convert users to students used for the SA algorithm
	// if any of the users cannot be converted, return an error
	// in practice this error should never happen
	for i, user := range users {
		cu, err := convertUserToStudent(user)

		if err != nil {
			return nil, err
		}

		convertedUsers[i] = cu
	}

	groups := DistributeStudents(convertedUsers, teamSizes)

	randomSource := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 1 = permutations size (length of the reverse movements slice)
	// false = no debug
	res := SA(groups, randomSource, saIterations, 1, saTemperature, false, saWeightSatisfaction, saWeightN, saWeightE, saWeightC)

	return &res, nil
}

// DistributeStudents distributes the students into teams of the given sizes
// Each team gets a student with high conscientiousness (C) as the first student and locks it in place
func DistributeStudents(students []Student, teamSizes []int) Teams {
	teams := make([]Team, len(teamSizes))

	// first sort student by their conscientiousness
	sort.Slice(students, func(i, j int) bool {
		return students[i].C > students[j].C
	})

	for i := 0; i < len(teamSizes); i++ {
		// Create the teams
		teams[i].Students = make([]Student, teamSizes[i])

		// Assign the the first N (len(teams)) most conscientious students to the teams
		// lock them in place to prevent them from being moved
		teams[i].Students[0] = students[i]
		teams[i].Students[0].Locked = true
	}

	studentIdx := len(teamSizes)

	for i := 0; i < len(teamSizes); i++ {
		for j := 1; j < teamSizes[i]; j++ {
			// Assign the rest of the students to the teams
			teams[i].Students[j] = students[studentIdx]
			studentIdx++
		}
	}

	return Teams(teams)
}
