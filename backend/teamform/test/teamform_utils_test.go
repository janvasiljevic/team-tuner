package teamform_test

import (
	"fmt"
	"jv/team-tone-tuner/teamform"
	"math/rand"
)

// ------------------- Helper functions -------------------

func generateRandomStudent(r *rand.Rand, name string, locked bool) teamform.Student {
	return teamform.Student{
		GithubNick: name,
		O:          r.Float64(),
		C:          r.Float64(),
		E:          r.Float64(),
		A:          r.Float64(),
		N:          r.Float64(),
		Locked:     locked,
	}
}

// generateRandomTeams generates a slice of teams with random students
// the first student of each team is locked (cannot be moved)
func generateRandomTeams(numOfGroups int, groupSize int, r *rand.Rand) teamform.Teams {

	teamsSlice := make([]teamform.Team, numOfGroups)

	for i := 0; i < numOfGroups; i++ {
		teamsSlice[i].Students = make([]teamform.Student, groupSize)
		for j := 0; j < groupSize; j++ {

			locked := false

			if j == 0 {
				locked = true
			}

			teamsSlice[i].Students[j] = generateRandomStudent(r, fmt.Sprintf("s-%d-%d", i, j), locked)

			teamsSlice[i].Dirty = true
		}
	}

	groups := teamform.Teams(teamsSlice)

	return groups
}

func getDefaultTestingTeams(r *rand.Rand) teamform.Teams {
	teamSlice := []teamform.Team{
		{
			Students: []teamform.Student{
				generateRandomStudent(r, "alice", true),
				generateRandomStudent(r, "bob", false),
				generateRandomStudent(r, "charlie", false),
			},
			Dirty: true,
		},
		{
			Students: []teamform.Student{
				generateRandomStudent(r, "david", true),
				generateRandomStudent(r, "eve", false),
				generateRandomStudent(r, "frank", false),
			},
			Dirty: true,
		},
		{
			Students: []teamform.Student{
				generateRandomStudent(r, "george", true),
				generateRandomStudent(r, "henry", false),
				generateRandomStudent(r, "isaac", false),
			},
			Dirty: true,
		},
	}

	teams := teamform.Teams(teamSlice)

	return teams
}
