package main

import (
	bigNoteApi "github.com/DiamondDmitriy/big-note-api"
	"github.com/DiamondDmitriy/big-note-api/config"
	"github.com/DiamondDmitriy/big-note-api/database"
	"github.com/DiamondDmitriy/big-note-api/internal/app/http/handler"
	"github.com/DiamondDmitriy/big-note-api/internal/app/http/route"
	"github.com/DiamondDmitriy/big-note-api/internal/core/service"
	"github.com/DiamondDmitriy/big-note-api/internal/repository"
	"log"
)

func main() {
	cnf, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error initializing config: %v", err)
	}

	db, _err := database.NewDB(&cnf.DB)
	if _err != nil {
		log.Fatalf("Error initializing database: %v", _err)
	}

	repositories := repository.NewRepositories(db)
	services := service.NewServices(repositories, cnf)
	controllers := handler.NewControllers(repositories, services)

	routes := &route.Route{
		Config:     cnf,
		Controller: controllers,
	}
	srv := new(bigNoteApi.Server)

	if err := srv.Run(cnf.HTTP.Port, routes.Init()); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
