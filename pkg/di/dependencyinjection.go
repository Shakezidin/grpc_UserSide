package di

import (
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
	adminhandler := handler.NewAdminHandler(userServ)
	server.GrpcServerRun(config, userhandler, adminhandler)
}
