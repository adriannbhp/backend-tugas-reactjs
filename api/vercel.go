package api

import (
	"backend-tugas-reactjs/config"
	"backend-tugas-reactjs/docs"
	"backend-tugas-reactjs/routes"
	"backend-tugas-reactjs/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	App *gin.Engine
)

func init() {
	App = gin.New()

	environment := utils.Getenv("ENVIRONMENT", "development")

	if environment == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	docs.SwaggerInfo.Title = "Book REST API"
	docs.SwaggerInfo.Description = "This is REST API Book."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = utils.Getenv("HOST", "localhost:8080")
	if environment == "development" {
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
	} else {
		docs.SwaggerInfo.Schemes = []string{"https"}
	}
	db := config.ConnectDataBase()

	routes.SetupRouter(db, App)
}

// Entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	App.ServeHTTP(w, r)
}

// func Handler(w http.ResponseWriter, rq *http.Request) {
// 	// for load godotenv
// 	// for env
// 	environment := utils.Getenv("ENVIRONMENT", "development")

// 	if environment == "development" {
// 		err := godotenv.Load()
// 		if err != nil {
// 			log.Fatal("Error loading .env file")
// 		}
// 	}

// 	//programmatically set swagger info
// 	docs.SwaggerInfo.Title = "Movie REST API"
// 	docs.SwaggerInfo.Description = "This is REST API Movie."
// 	docs.SwaggerInfo.Version = "1.0"
// 	docs.SwaggerInfo.Host = utils.Getenv("HOST", "localhost:8080")
// 	if environment == "development" {
// 		docs.SwaggerInfo.Schemes = []string{"http", "https"}
// 	} else {
// 		docs.SwaggerInfo.Schemes = []string{"https"}
// 	}

// 	db := config.ConnectDataBase()
// 	sqlDB, _ := db.DB()
// 	defer sqlDB.Close()

// 	r := routes.SetupRouter(db)
// 	r.Run()
// }
