package controllers

import (
	"strconv"
	"yquiz_back/internal/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userForm models.UserForm

	if err := c.ShouldBind(&userForm); err != nil {
		c.JSON(400, gin.H{
			"message": "Données de formulaire invalides",
			"error":   err.Error(),
		})
		return
	}

	user, err := userForm.ToUser()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Erreur lors de la création de l'utilisateur",
			"error":   err.Error(),
		})
		return
	}

	/*
		TODO:
		Requete dans la base de données pour créer l'utilisateur
	*/

	c.JSON(200, gin.H{
		"message": "User created",
		"user":    user,
	})

}

func GetUser(c *gin.Context) {

	users := []models.User{
		{
			ID:        1,
			Email:     "exemple1@domaine.com",
			Password:  "motdepassehashé1",
			FirstName: "Prénom1",
			LastName:  "Nom1",
			Role:      "student",
			ClassID:   nil,
		},
		{
			ID:        2,
			Email:     "exemple2@domaine.com",
			Password:  "motdepassehashé2",
			FirstName: "Prénom2",
			LastName:  "Nom2",
			Role:      "teacher",
			ClassID:   nil,
		},
		{
			ID:        3,
			Email:     "exemple3@domaine.com",
			Password:  "motdepassehashé3",
			FirstName: "Prénom3",
			LastName:  "Nom3",
			Role:      "admin",
			ClassID:   nil,
		},
	}

	id := c.Param("id")

	if id == "" {
		c.JSON(400, gin.H{
			"message": "ID invalide",
		})
		return
	}

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "ID invalide",
		})
		return
	}

	for _, user := range users {
		if user.ID == uint(idUint) {
			c.JSON(200, gin.H{
				"message": "User retrieved",
				"user":    user,
			})
		}
	}

}
