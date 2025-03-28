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

	teacherID := controllers.GetTeacherID(c)
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

func GetQuizzes(c *gin.Context) {
	teacherID := controllers.GetTeacherID(c)
	if teacherID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Accès refusé"})
		return
	}

	search, sortBy, page, limit, err := controllers.ParseQuizQueryParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	quizzes, total, err := controllers.GetQuizzes(teacherID, search, sortBy, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erreur serveur"})
		return
	}

	var res []gin.H
	for _, q := range quizzes {
		res = append(res, gin.H{
			"id":          q.ID,
			"title":       q.Title,
			"description": q.Description,
			"duration":    q.Duration,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
		"pagination": models.Pagination{
			CurrentPage: page,
			TotalItems:  int(total),
			TotalPages:  (int(total) + limit - 1) / limit,
		},
	})
}
