package cmd

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Environment map[string]string
var DbConn *gorm.DB
var Log *logrus.Logger
var GinSv *gin.Engine

func StartApp() {
	initEnv()
	initLog()
	initDatabase()
	initGin()
}

func initEnv() {
	envVars, err := godotenv.Read()

	if err != nil {
		fmt.Println(err.Error())
	}

	Environment = envVars
	fmt.Printf("Environment variables loaded in %s profile. \n", strings.ToUpper(Environment["CURRENT_ENV"]))
}

func initLog() {
	log := logrus.New()

	lvl, _ := logrus.ParseLevel(Environment["LOG_LEVEL"])
	log.SetLevel(lvl)

	log.SetFormatter(&logrus.TextFormatter{DisableColors: true})

	Log.Info("Logger initialized")

	Log = log
}

func initDatabase() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		Environment["DB_HOST"],
		Environment["DB_PORT"],
		Environment["DB_USERNAME"],
		Environment["DB_PASSWORD"],
		Environment["DB_NAME"],
	)

	dbc, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		Log.Error(err.Error())
	}

	DbConn = dbc
}

func initGin() {
	ginMode := gin.DebugMode
	switch Environment["CURRENT_ENV"] {
	case "DEV":
		ginMode = gin.ReleaseMode
	case "TEST":
		ginMode = gin.TestMode
	case "PROD":
		ginMode = gin.ReleaseMode
	}

	gin.SetMode(ginMode)

	r := gin.New()

	requestLogging := func(ctx *gin.Context) {
		ctx.Next()

		lgf := logrus.Fields{
			"type":   "HTTP Request",
			"uri":    ctx.Request.RequestURI,
			"method": ctx.Request.Method,
			"status": ctx.Writer.Status(),
		}

		Log.WithFields(lgf).Info()

		ctx.Next()
	}

	if ginMode == gin.DebugMode || ginMode == gin.TestMode {
		r.Use(requestLogging)
	}

	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	GinSv = r
	r.Run(":" + Environment["SERVER_PORT"])

}
