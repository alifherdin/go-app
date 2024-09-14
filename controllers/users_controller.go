package controllers

import (
	"go-api/dtos/userdtos"
	"go-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(e *gin.Engine, svc *services.UserService) {
	ctl := UserController{svc}

	public := e.Group("/users")
	public.GET("/signup", ctl.signup)
}

func (uc *UserController) signup(c *gin.Context) {
	req := userdtos.CreateUserRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// res, err := uc.UserService.CreateUser(req)

	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusCreated, gin.H{"data": req})
}
