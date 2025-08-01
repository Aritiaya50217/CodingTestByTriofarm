package main

import (
	"log"
	"os"

	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/config"
	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/delivery/http"
	database "github.com/Aritiaya50217/CodingTestByTriofarm/internal/infrastructure/db"
	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/repository"
	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	config.SetTimeZone()

	db, err := database.InitDB(cfg.DSN)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	// CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	api := r.Group("/v1/api")

	// topic
	topicRepo := repository.NewTopicRepository(db)
	topicUsecase := usecase.NewTopicUsecase(topicRepo)
	http.NewTopicHandler(r, topicUsecase, api)

	// medicine
	medicineRepo := repository.NewMedicineRepository(db)
	medicineUsecase := usecase.NewMedicineUsecase(medicineRepo)
	http.NewMedicineHandler(r, medicineUsecase, api)

	// vitamin
	vitaminRepo := repository.NewVitaminRepository(db)
	vitaminUsecase := usecase.NewVitaminUsecase(vitaminRepo)
	http.NewVitaminHandler(r, vitaminUsecase, api)

	// microorganism
	microorganismRepo := repository.NewMicroorganismRepository(db)
	microorganismUsecase := usecase.NewMicroorganismUsecase(microorganismRepo)
	http.NewMicroorganismHandler(r, microorganismUsecase, api)

	// brand
	brandRepo := repository.NewBrandRepository(db)
	brandUsecase := usecase.NewBrandUsecase(brandRepo)
	http.NewBrandHandler(r, brandUsecase, api)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
