package main

import (
	"context"
	"flag"
	"jv/team-tone-tuner/model"
	"jv/team-tone-tuner/model/user"

	_ "github.com/lib/pq"
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

	client.User.Create().SetGithubUsername("janvasiljevic").SetRole(user.RoleAdmin).SaveX(ctx)
}
