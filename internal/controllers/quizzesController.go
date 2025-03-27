package controllers

import (
	"errors"
	"yquiz_back/internal/database"
	"yquiz_back/internal/models"
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
