package internal

import (
	"yquiz_back/internal/database"

	"github.com/gin-gonic/gin"
)

func Init_server() {
	database.InitDatabase()

	router := gin.Default()
	router.GET("/", Home)
	router.Run(":8080")

}
