package handler

import (
	"net/http"
	"yquiz_back/internal/controllers"
	"yquiz_back/internal/models"

	"github.com/gin-gonic/gin"
)

func CreateQuiz(c *gin.Context) {
	var quizForm models.QuizzForm
	if err := c.ShouldBindJSON(&quizForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Données invalides"})
		return
	}

	teacherID := controllers.CheckUserRole(c)
	if teacherID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Accès refusé"})
		return
	}

	quiz, err := controllers.CreateQuiz(quizForm, teacherID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erreur lors de la création du quiz",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"quiz_id":     quiz.ID,
			"title":       quiz.Title,
			"description": quiz.Description,
			"duration":    quiz.Duration,
		},
	})
}
