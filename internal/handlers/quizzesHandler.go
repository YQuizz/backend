package handler

import (
	"net/http"
	"yquiz_back/internal/controllers"
	"yquiz_back/internal/database"
	"yquiz_back/internal/models"

	"github.com/gin-gonic/gin"
)

// Faire une transaction pour créer le quiz et les questions uniquement si tout est valide
func CreateQuiz(c *gin.Context) {

	var quizForm models.QuizzForm
	if err := c.ShouldBindJSON(&quizForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Données invalides",
			"error":   err.Error(),
		})
		return
	}

	teacherID := controllers.CheckUserRole(c)
	if teacherID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	quiz := models.Quiz{
		Title:       quizForm.Title,
		Description: quizForm.Description,
		Duration:    quizForm.Duration,
		TeacherID:   teacherID,
	}

	if err := database.DB.Create(&quiz).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erreur lors de la création du quiz",
			"error":   err.Error(),
		})
		return
	}

	for _, questionForm := range quizForm.Questions {
		question := models.Question{
			QuizID: quiz.ID,
			Text:   questionForm.Text,
			Type:   questionForm.Type,
		}

		if err := database.DB.Create(&question).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Erreur lors de la création des questions",
				"error":   err.Error(),
			})
			return
		}

		for _, answerForm := range questionForm.Answers {
			answer := models.Answer{
				QuestionID: question.ID,
				Text:       answerForm.Text,
				IsCorrect:  answerForm.IsCorrect,
			}

			if err := database.DB.Create(&answer).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Erreur lors de la création des réponses",
					"error":   err.Error(),
				})
				return
			}
		}
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
