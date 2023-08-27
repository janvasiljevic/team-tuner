package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	if err := entc.Generate("./schema", &gen.Config{
		Target:  "./model",
		Package: "jv/team-tone-tuner/model",
		Features: []gen.Feature{
			gen.FeatureUpsert,
		},
	}); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
