package cmd

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var EnvironmentVars map[string]string
var DbConn *gorm.DB
var log *zap.Logger

func StartApp() {
	initEnv()
	initLog()
	initDatabase()
}

func initEnv() {
	envVars, err := godotenv.Read()

	if err != nil {
		fmt.Println(err.Error())
	}

	EnvironmentVars = envVars
	fmt.Printf("Environment variables loaded in %s profile.", strings.ToUpper(EnvironmentVars["CURRENT_ENV"]))
}

func initLog() {
	config := zap.NewDevelopmentConfig() // why DevelopmentConfig? because colors.
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.DisableStacktrace = true

	lvl := zapcore.InfoLevel
	switch EnvironmentVars["LOG_LEVEL"] {
	case "DEBUG":
		lvl = zapcore.DebugLevel
	case "INFO":
		lvl = zapcore.InfoLevel
	case "WARN":
		lvl = zapcore.WarnLevel
	case "ERROR":
		lvl = zapcore.ErrorLevel
	case "DPANIC":
		lvl = zapcore.DPanicLevel
	case "PANIC":
		lvl = zapcore.PanicLevel
	case "FATAL":
		lvl = zapcore.FatalLevel
	}

	config.Level.SetLevel(lvl)

	logger, _ := config.Build()
	log = logger
}

func initDatabase() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		EnvironmentVars["DB_HOST"],
		EnvironmentVars["DB_PORT"],
		EnvironmentVars["DB_USERNAME"],
		EnvironmentVars["DB_PASSWORD"],
		EnvironmentVars["DB_NAME"],
	)

	dbc, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Error(err.Error())
	}

	DbConn = dbc
}
