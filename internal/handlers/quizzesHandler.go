package handler

import (
	"net/http"
	"yquiz_back/internal/controllers"
	"yquiz_back/internal/database"
	"yquiz_back/internal/models"

	"github.com/gin-gonic/gin"
)

func CreateQuiz(c *gin.Context) {

	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if error := tx.Error; error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erreur lors de la connexion",
		})
		return
	}

	var quizForm models.QuizzForm
	if err := c.ShouldBindJSON(&quizForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Données invalides",
		})
		return
	}

	teacherID := controllers.CheckUserRole(c)
	if teacherID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Accès refusé",
		})
		return
	}

	quiz := models.Quiz{
		Title:       quizForm.Title,
		Description: quizForm.Description,
		Duration:    quizForm.Duration,
		TeacherID:   teacherID,
	}

	if err := tx.Create(&quiz).Error; err != nil {
		tx.Rollback()
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

		if err := tx.Create(&question).Error; err != nil {
			tx.Rollback()
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

			if err := tx.Create(&answer).Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Erreur lors de la création des réponses",
					"error":   err.Error(),
				})
				return
			}
		}
	}

	tx.Commit()

	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"quiz_id":     quiz.ID,
			"title":       quiz.Title,
			"description": quiz.Description,
			"duration":    quiz.Duration,
		},
	})
}
