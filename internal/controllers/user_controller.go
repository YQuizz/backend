package controllers

import (
	"yquiz_back/internal/database"
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

	result := database.DB.Create(user)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Erreur lors de la création de l'utilisateur",
			"error":   result.Error.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "User created",
		"user":    user,
	})

}

func GetUser(c *gin.Context) {

	id := c.Param("id")

	if id == "" {
		c.JSON(400, gin.H{
			"message": "ID invalide",
		})
		return
	}

	var user models.User
	result := database.DB.First(&user, "id = ?", id)
	if result.Error != nil {
		c.JSON(404, gin.H{
			"message": "Utilisateur non trouvé",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User retrieved",
		"user":    user,
	})
}
