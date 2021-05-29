package main

import (
	"github.com/SupratickDey/go_gin_backend/api"
	"github.com/SupratickDey/go_gin_backend/database"
	"github.com/SupratickDey/go_gin_backend/lib/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// load .env environment variables
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// initializes database
	db, _ := database.Initilize()

	port := os.Getenv("PORT")
	app := gin.Default() // create gin app
	app.Use(database.Inject(db))
	app.Use(middlewares.JWTMiddleware())
	api.ApplyRoutes(app) // apply api router
	log.Fatal(app.Run(":" + port))  // listen to given port
}

