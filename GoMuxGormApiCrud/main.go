package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/melardev/GoMuxGormApiCrud/infrastructure"
	"github.com/melardev/GoMuxGormApiCrud/models"
	"github.com/melardev/GoMuxGormApiCrud/routes"
	"github.com/melardev/GoMuxGormApiCrud/seeds"
	"net/http"
	"os"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Todo{})
}

func main() {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
		os.Exit(0)
	}

	database := infrastructure.OpenDbConnection()
	defer database.Close()
	migrate(database)
	seeds.Seed(database)

	routes.RegisterRoutes()

	// corsMiddleware := handlers.CORS()
	// It is just to make it more readable, corsMiddleware(routes) is the same as handlers.CORS()(routes)
	if err := http.ListenAndServe(":8080", handlers.CORS()(routes.Router)); err != nil {
		println(err.Error())
	}
}
