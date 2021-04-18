package main

import (
	"log"
	"os"

	"github.com/niltonkummer/desafio-conductor/app/config"

	"gorm.io/driver/sqlite"

	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/niltonkummer/desafio-conductor/app"
)

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

// @info
// @version: 0.1.0
// @title Desafio Conductor

// @host warm-bastion-37111.herokuapp.com
// @schemes https http
// @BasePath /conductor/v1/api
func main() {

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	dbDialect := os.Getenv("DB")
	if dbDialect == "" {
		log.Fatal("$DB must be set")
	}

	app.NewApplication(&config.Config{
		HttpListen: ":" + port,
		DBDialect:  sqlite.Open(dbDialect),
	}).StartServer()
}
