package utils

import (
	"github.com/gin-gonic/gin"
)

func BindJsonRequest(ctx *gin.Context, obj any) (any, error) {
	err := ctx.BindJSON(&obj)
	if err != nil {
		return nil, err
	}
	return obj, err
}
