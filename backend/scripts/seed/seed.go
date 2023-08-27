package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"

	"jv/team-tone-tuner/model"
	"jv/team-tone-tuner/model/bfiquestion"
	"jv/team-tone-tuner/schema"
	"log"
	"math/rand"
	"os"
	"time"

	_ "github.com/lib/pq"

	"gonum.org/v1/gonum/stat/distuv"
)

type VprasanjaJson struct {
	Vprasanja []Vprasanja `json:"vprasanja"`
}

type Vprasanja struct {
	N            int64        `json:"n"`
	QuestionSLO  string       `json:"Question (SLO)"`
	QuestionENG  string       `json:"Question (ENG)"`
	FacetSLO     string       `json:"Facet (SLO)"`
	FacetENG     string       `json:"Facet (ENG)"`
	DimensionENG DimensionENG `json:"Dimension (ENG)"`
	Influence    Influence    `json:"Influence"`
	Alpha        float64      `json:"Alpha"`
	Bfi          int64        `json:"BFI"`
}

type DimensionENG string

const (
	Agreeableness        DimensionENG = "agreeableness"
	Conscientiousness    DimensionENG = "conscientiousness"
	Extraversion         DimensionENG = "extraversion"
	Neuroticism          DimensionENG = "neuroticism"
	OpennessToExperience DimensionENG = "openness to experience"
)

type Influence string

const (
	Negative Influence = "negative"
	Positive Influence = "positive"
)

func main() {

	var dsnFlag string

	d := "host=localhost port=5432 user=dev dbname=dev password=12345678 sslmode=disable"
	flag.StringVar(&dsnFlag, "dsn", d, "Postgres DSN")

	client, err := model.Open("postgres", dsnFlag)

	if err != nil {
		panic(err)
	}

	defer client.Close()

	ctx := context.Background()

	client.Course.Delete().ExecX(ctx)
	client.BfiReport.Delete().ExecX(ctx)
	client.User.Delete().ExecX(ctx)
	client.BfiQuestion.Delete().ExecX(ctx)

	// read the file into the json
	var vprasanjaJson VprasanjaJson

	jsonFile, err := os.Open("./resources/seed/vprasanja.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		fmt.Println(err)
	}

	_ = json.Unmarshal(byteValue, &vprasanjaJson)

	// Create all questions
	bulkBfiQuestion := make([]*model.BfiQuestionCreate, len(vprasanjaJson.Vprasanja))

	for i, vprasanje := range vprasanjaJson.Vprasanja {
		bulkBfiQuestion[i] = client.BfiQuestion.Create().
			SetAlpha(vprasanje.Alpha).
			SetQuestiono(vprasanje.QuestionENG).
			SetFacet(vprasanje.FacetENG)

		switch vprasanje.DimensionENG {
		case Agreeableness:
			bulkBfiQuestion[i] = bulkBfiQuestion[i].SetDimension(bfiquestion.DimensionAgreeableness)
		case Conscientiousness:
			bulkBfiQuestion[i] = bulkBfiQuestion[i].SetDimension(bfiquestion.DimensionConscientiousness)
		case Extraversion:
			bulkBfiQuestion[i] = bulkBfiQuestion[i].SetDimension(bfiquestion.DimensionExtraversion)
		case Neuroticism:
			bulkBfiQuestion[i] = bulkBfiQuestion[i].SetDimension(bfiquestion.DimensionNeuroticism)
		case OpennessToExperience:
			bulkBfiQuestion[i] = bulkBfiQuestion[i].SetDimension(bfiquestion.DimensionOpenness)
		}

		switch vprasanje.Influence {
		case Negative:
			bulkBfiQuestion[i] = bulkBfiQuestion[i].SetInfluence(bfiquestion.InfluenceNegative)
		case Positive:
			bulkBfiQuestion[i] = bulkBfiQuestion[i].SetInfluence(bfiquestion.InfluencePositive)
		}
	}

	client.BfiQuestion.CreateBulk(bulkBfiQuestion...).SaveX(ctx)

	log.Printf("Created %d questions", len(bulkBfiQuestion))

	// Create one course
	course := client.Course.Create().
		SetName("Software Engineering").
		SetCode("SENG1112").
		SetColour("#ff0000").
		SaveX(ctx)

	log.Printf("Created course with name %s", course.Name)

	numOfStudents := 150

	bulkStudent := make([]*model.UserCreate, numOfStudents)

	for i := 0; i < numOfStudents; i++ {
		bulkStudent[i] = client.User.Create().
			SetGithubUsername(fmt.Sprintf("%d studentl", i)).
			AddCourses(course)
	}

	students := client.User.CreateBulk(bulkStudent...).SaveX(ctx)

	log.Printf("Created %d students", len(students))

	bulkBfiReport := make([]*model.BfiReportCreate, numOfStudents)

	mus := []float64{0.55, 0.45, 0.52, 0.49, 0.65}

	dists := make([]distuv.Normal, 5)

	for i := 0; i < 5; i++ {
		dists[i] = distuv.Normal{
			Mu:    mus[i],
			Sigma: 0.1,
		}
	}

	for i := 0; i < numOfStudents; i++ {

		// Create a random date between today and 7 days ago
		n := rand.Intn(8) // Generate a random integer between 0 and 7
		date := time.Now().AddDate(0, 0, -n)

		bulkBfiReport[i] = client.BfiReport.Create().
			SetStudentID(students[i].ID).
			SetAgreeableness(createRandomReport(dists[0])).
			SetConscientiousness(createRandomReport(dists[1])).
			SetExtraversion(createRandomReport(dists[2])).
			SetNeuroticism(createRandomReport(dists[3])).
			SetOpenness(createRandomReport(dists[4])).
			SetCreatedAt(date)
	}

	client.BfiReport.CreateBulk(bulkBfiReport...).SaveX(ctx)

	log.Printf("Created %d BFI reports", len(students))
}

func createRandomReport(d distuv.Normal) schema.BfiReportItem {
	return schema.BfiReportItem{
		PointsMin:        0,
		PointsMax:        30,
		PointsNormalized: d.Rand(),
		PointsSum:        15,
	}
}
