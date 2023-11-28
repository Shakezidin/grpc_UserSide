package di

import (
	"log"

	"github.com/shakezidin/config"
	"github.com/shakezidin/pkg/db"
	"github.com/shakezidin/pkg/handler"
	repo "github.com/shakezidin/pkg/repository"
	"github.com/shakezidin/pkg/server"
	serv "github.com/shakezidin/pkg/service"
)

func Init() {
	config := config.LoadConfig()
	db := db.NewDBConnect(config)
	userRepo := repo.NewUserRepository(db)
	userServ := serv.NewUserService(userRepo)
	userhandler := handler.NewUserHandler(userServ)
	err := server.NewGrpcServer(config, userhandler)
	if err != nil {
		log.Fatalf("something went wrong", err)
	}
}
