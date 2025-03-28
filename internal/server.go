package internal

import (
	"log"
	"yquiz_back/internal/database"
	handler "yquiz_back/internal/handlers"

	"github.com/gin-gonic/gin"
)

func Init_server() {
	_, err := database.GetDB()
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	router := gin.Default()
	initRoutes(router)
	router.Run(":8080")
}

func initRoutes(router *gin.Engine) {

	api := router.Group("/api")
	{
		/*
			Users API routes
		*/
		api.POST("/login", handler.Login)

		/*
			Quizzes API routes
		*/
		api.POST("/quizzes", handler.CreateQuiz)
		api.GET("/quizzes", handler.GetQuizzes)

	}
}
