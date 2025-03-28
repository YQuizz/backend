package controllers

import (
	"errors"
	"fmt"
	"strings"
	"yquiz_back/internal/database"
	"yquiz_back/internal/models"
	"yquiz_back/internal/pkg"

	"github.com/gin-gonic/gin"
)

func CreateQuiz(quizForm models.QuizzForm, teacherID uint) (models.Quiz, error) {
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return models.Quiz{}, errors.New("Erreur lors de la connexion")
	}

	quiz := models.Quiz{
		Title:       quizForm.Title,
		Description: quizForm.Description,
		Duration:    quizForm.Duration,
		TeacherID:   teacherID,
	}

	if err := tx.Create(&quiz).Error; err != nil {
		tx.Rollback()
		return models.Quiz{}, err
	}

	for _, questionForm := range quizForm.Questions {
		question := models.Question{
			QuizID: quiz.ID,
			Text:   questionForm.Text,
			Type:   questionForm.Type,
		}
		if err := tx.Create(&question).Error; err != nil {
			tx.Rollback()
			return models.Quiz{}, err
		}

		for _, answerForm := range questionForm.Answers {
			answer := models.Answer{
				QuestionID: question.ID,
				Text:       answerForm.Text,
				IsCorrect:  answerForm.IsCorrect,
			}
			if err := tx.Create(&answer).Error; err != nil {
				tx.Rollback()
				return models.Quiz{}, err
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		return models.Quiz{}, err
	}

	return quiz, nil
}

func GetQuizzes(teacherID uint, search, sortBy string, page, limit int) ([]models.Quiz, int64, error) {
	query := database.DB.Where("teacher_id = ?", teacherID)

	if search != "" {
		query = query.Where("LOWER(title) LIKE ?", "%"+strings.ToLower(search)+"%")
	}

	switch sortBy {
	case "oldest":
		query = query.Order("id ASC")
	case "duration_asc":
		query = query.Order("duration ASC")
	case "duration_desc":
		query = query.Order("duration DESC")
	default:
		query = query.Order("id DESC")
	}

	var total int64
	query.Model(&models.Quiz{}).Count(&total)

	offset := (page - 1) * limit
	var quizzes []models.Quiz
	err := query.Offset(offset).Limit(limit).Find(&quizzes).Error
	return quizzes, total, err
}

func ParseQuizQueryParams(c *gin.Context) (search string, sortBy string, page int, limit int, err error) {
	search = c.Query("search")
	rawSortBy := c.DefaultQuery("sort_by", "recent")
	validSorts := map[string]bool{
		"recent":        true,
		"oldest":        true,
		"duration_asc":  true,
		"duration_desc": true,
	}
	if !validSorts[rawSortBy] {
		err = fmt.Errorf("tri invalide")
		return
	}
	sortBy = rawSortBy

	page, limit = pkg.GetPaginationParams(c)
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	if limit > 100 {
		err = fmt.Errorf("limite max = 100")
		return
	}

	return
}
