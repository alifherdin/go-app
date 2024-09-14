package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Environment map[string]string
var DbConn *gorm.DB
var Log *logrus.Logger
var GinSv *gin.Engine
