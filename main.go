package main

import (
	"log"
	"task-5-pbi-btpns-arufhakim/database"
	"task-5-pbi-btpns-arufhakim/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	if err := database.Open(); err != nil {
		log.Fatal(err)
	}

	app := router.Start()
	log.Fatal(app.Run("127.0.0.1:8080"))
}
