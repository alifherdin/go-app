package controllers

import (
	"fmt"
	"go-api/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(e *gin.Engine, svc *services.UserService) *gin.Engine {
	ct := &UserController{svc}

	usersPublic := e.Group("/users")
	usersPublic.POST("/login", func(ctx *gin.Context) {})

	fmt.Printf("ct: %v\n", ct)

	return e
}
