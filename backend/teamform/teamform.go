package teamform

import (
	"fmt"
	"jv/team-tone-tuner/utils"
	"math/rand"
	"strings"

	"github.com/google/uuid"
)

// Grouper is the package which contains the logic for grouping the
// students into groups.

type Student struct {
	GithubNick    string
	Id            uuid.UUID
	O, C, E, A, N float64
	Locked        bool
}

type Team struct {
	Students   []Student
	S, C, E, N float64
	Dirty      bool
}

type Teams []Team

type MovementItem struct {
	FromStudentIndex,
	FromGroupIndex,
	ToStudentIndex,
	ToGroupIndex int
}

// NewMovementsSlice creates a slice of MovementItems with the given size
func NewMovementsSlice(size int) []MovementItem {
	return make([]MovementItem, size)
}

// AddOCEANToDenseMatrix adds the OCEAN values of the student to the dense matrix
// for compute in the XGBoost model
func (s Student) AddOCEANToDenseMatrix(dm *[]float64) {
	*dm = append(*dm, s.O, s.C, s.E, s.A, s.N)
}

func (teams Teams) ShuffleStudents(random *rand.Rand, reverseMovements []MovementItem) {

	steps := len(reverseMovements)

	if steps == 0 {
		panic("shuffle size cannot be 0")
	}

	for i := steps - 1; i >= 0; i-- {
		teamIdx1, teamIdx2 := randomPairWithoutDuplicates(random, 0, len(teams))

		students1 := teams[teamIdx1].Students
		students2 := teams[teamIdx2].Students

		studentIdx1 := teams[teamIdx1].PickARandomNonLockedStudentIndex()
		studentIdx2 := teams[teamIdx2].PickARandomNonLockedStudentIndex()

		// Swap students between the groups
		students1[studentIdx1], students2[studentIdx2] = students2[studentIdx2], students1[studentIdx1]

		// Update the reverse movements
		reverseMovements[i] = MovementItem{
			FromStudentIndex: studentIdx1,
			FromGroupIndex:   teamIdx1,
			ToStudentIndex:   studentIdx2,
			ToGroupIndex:     teamIdx2,
		}

		// Mark the swapped teams as dirty - to recalculate team satisfaction on the next iteration
		teams[teamIdx1].Dirty = true
		teams[teamIdx2].Dirty = true
	}
}

// ReverseStudentShuffles reverse the permutation of the Teams
func (teams Teams) ReverseStudentShuffles(movements []MovementItem) {
	for _, m := range movements {
		teams[m.FromGroupIndex].Students[m.FromStudentIndex],
			teams[m.ToGroupIndex].Students[m.ToStudentIndex] =
			teams[m.ToGroupIndex].Students[m.ToStudentIndex],
			teams[m.FromGroupIndex].Students[m.FromStudentIndex]

		teams[m.FromGroupIndex].Dirty = false
		teams[m.ToGroupIndex].Dirty = false
	}
}

// ToString returns a string representation of the Teams
func (teams Teams) ToString() string {
	builder := ""

	for _, team := range teams {
		studentInfo := utils.Map(team.Students, func(s Student) string {
			return fmt.Sprintf("%10s [%.2f, %.2f, %.2f, %.2f, %.2f]", s.GithubNick, s.O, s.C, s.E, s.A, s.N)
		})

		builder += fmt.Sprintf("{%s}\n", strings.Join(studentInfo, " "))
	}

	return builder
}

// Reset clears the 'cache' of the team - the scores (and the auxiliary dirty flag)
func (team *Team) Reset() {
	team.S, team.C, team.E, team.N = 0, 0, 0, 0
	team.Dirty = false
}

// PickARandomNonLockedStudentIndex picks a random student index from the team
// It relies on the fact, that only the first student is locked
// Just for sanity reasons we put this into a loop
func (team *Team) PickARandomNonLockedStudentIndex() int {
	for {
		studentIdx := rand.Intn(len(team.Students)-1) + 1

		if !team.Students[studentIdx].Locked {
			return studentIdx
		}
	}
}
