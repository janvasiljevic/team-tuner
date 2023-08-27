package main

import (
	"fmt"
	"os"

	"jv/team-tone-tuner/api"
	"jv/team-tone-tuner/config"
	"jv/team-tone-tuner/model"
	"jv/team-tone-tuner/router"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	_ "jv/team-tone-tuner/docs"

	_ "github.com/lib/pq"
	echoSwagger "github.com/swaggo/echo-swagger"
)

//	@title			Team tuner API
//	@version		1.0
//	@description	Team tuner API documentation
//	@title			TT API

//	@BasePath	/api

//	@schemes	http https
//	@produce	application/json
//	@consumes	application/json
func main() {
	// https://github.com/xesina/golang-echo-realworld-example-app

	cnf, err := config.NewEnv()

	if err != nil {
		log.Error().Msg(fmt.Sprintf("failed loading config: %v", err))

		os.Exit(-1)
	}

	config.LoadedConfig = *cnf

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	r := router.New()
	r.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := r.Group("/api")

	// Instantiate the Ent client.
	client, err := model.Open("postgres", config.LoadedConfig.Database.ConnectionUrl)

	if err != nil {
		log.Error().Msg(fmt.Sprintf("failed opening connection to db: %v", err))

		os.Exit(-1)
	}

	defer client.Close()

	api := api.New(client)

	api.RegisterRoutes(v1)

	r.Logger.Fatal(r.Start(fmt.Sprintf(":%d", config.LoadedConfig.Server.Port)))
}
