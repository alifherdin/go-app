package controllers

import (
	"go-api/dtos/userdtos"
	"go-api/services"
	"go-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(e *gin.Engine, svc *services.UserService) {
	ctl := UserController{svc}

	public := e.Group("/users")
	public.POST("/signup", ctl.signup)
}

func (uc *UserController) signup(c *gin.Context) {
	req, err := utils.BindJsonRequest(c, &userdtos.CreateUserRequest{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := uc.UserService.CreateUser(req.(userdtos.CreateUserRequest))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": res})
}
