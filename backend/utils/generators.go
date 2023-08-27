package utils

import "math/rand"

var (
	colors     []string
	adjectives []string
	surnames   []string
)

func init() {
	colors = []string{"red", "blue", "green", "yellow", "orange", "pink", "brown", "purple", "gray", "white"}

	adjectives = []string{"adoring", "affectionate", "agitated", "amazing", "angry", "awesome", "beautiful", "blissful", "bold", "boring",
		"brave", "busy", "charming", "clever", "cool", "compassionate", "competent", "condescending", "confident", "crazy"}

	surnames = []string{"albattani", "allen", "almeida", "antonelli", "agnesi", "archimedes", "ardinghelli", "aryabhata", "austin", "babbage",
		"banach", "bardeen", "bartik", "bassi", "beaver", "bell", "benz", "bhabha", "bhaskara", "blackwell",
		"bohr", "booth", "borg", "bose", "bouman", "boyd", "brahmagupta", "brattain", "brown", "carson",
		"chandrasekhar", "chatelet", "chatterjee", "chebyshev", "cohen", "chaum", "clarke", "colden", "cori", "cray"}
}

func GenerateRandomDockerLikeName(rand *rand.Rand) string {
	// Generate a random combination of an adjective and a surname
	color := colors[rand.Intn(len(colors))]
	adjective := adjectives[rand.Intn(len(adjectives))]
	surname := surnames[rand.Intn(len(surnames))]

	return color + " " + adjective + " " + surname
}
