package internal

import (
	"log"
	"yquiz_back/internal/controllers"
	"yquiz_back/internal/database"

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
		api.POST("/login", controllers.Login)

		/*
			Quizzes API routes
		*/

	}
}
