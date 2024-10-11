package main

import (
	"github.com/minnowo/traggo-sync/client"
	"github.com/minnowo/traggo-sync/logging"
	"github.com/rs/zerolog/log"
)

const traggoUrl string = "http://localhost:3030"

func main() {

	logging.InitFromEnv()

	traggoClient, err := client.NewTraggoClient(traggoUrl)

	if err != nil {

		log.Fatal().Err(err)
	}

	traggoClient.LogVersion()

	err = traggoClient.Login("admin", "admin")

	if err != nil {

		log.Fatal().Err(err)
	}

}
