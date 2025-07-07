package main

import (
	bigNoteApi "github.com/DiamondDmitriy/big-note-api"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error initializing config: %v", err)
	}

	handlers := new(bigNoteApi.Handler)
	srv := new(bigNoteApi.Server)
	if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
