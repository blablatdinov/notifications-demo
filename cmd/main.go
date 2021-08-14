package main

import (
	"log"

	"github.com/blablatdinov/notifications-demo"
	"github.com/blablatdinov/notifications-demo/pkg/handler"
	"github.com/blablatdinov/notifications-demo/pkg/repository"
	"github.com/blablatdinov/notifications-demo/pkg/service"
)

func main() {
	db, err := repository.NewPostgresDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	services.Notifications.PingUsers()
	handlers := handler.NewHandler(services)
	srv := new(notifications.Server)
	if err := srv.Run(handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
