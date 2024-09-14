package cmd

import (
	"go-api/controllers"
	"go-api/repositories"
	"go-api/services"
)

func InitRoutes() {
	userSvc := services.NewUserService(repositories.NewUserRepository(DbConn))

	controllers.NewUserController(GinSv, userSvc)
}
