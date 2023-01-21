package main

import (
	"Maybe/src/modules"
	"Maybe/src/modules/server"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	app, version := "maybe", 1.0
	log.Println("running ", app, " at version ", version)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	err = modules.MongoInit()
	if err != nil {
		log.Fatal(err)
	}

	err = server.Init()
	if err != nil {
		log.Fatal(err)
	}
}
