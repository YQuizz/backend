package handler

import (
	"yquiz_back/internal/controllers"
	"yquiz_back/internal/database"
	"yquiz_back/internal/models"
	"yquiz_back/internal/pkg"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	/* {
	    "message": "User created",
	    "user": {
	        "id": 2,
	        "email": "mail@mail2.com",
	        "first_name": "ludo",
	        "last_name": "roux",
	        "role": "admin",
	        "class_id": null,
	        "Class": null,
	        "Quizzes": null,
	        "UserAnswers": null,
	        "MonitoringLogs": null
	    }
	} */

	var loginForm models.LoginForm

	if err := c.ShouldBind(&loginForm); err != nil {
		c.JSON(400, gin.H{
			"message": "Données de connexion invalides",
			"error":   err.Error(),
		})
		return
	}

	var user models.User
	result := database.DB.Where("email = ?", loginForm.Email).First(&user)
	if result.Error != nil {
		c.JSON(401, gin.H{
			"message": "Email ou mot de passe incorrect",
		})
		return
	}

	if !pkg.CheckPassword(loginForm.Password, user.Password) {
		c.JSON(401, gin.H{
			"message": "Email ou mot de passe incorrect",
		})
		return
	}

	token := controllers.CreateJWT(c, &user)

	c.JSON(200, gin.H{
		"data": gin.H{
			"user_id":    user.ID,
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"role":       user.Role,
			"class_id":   user.ClassID,
			"token":      token,
		},
	})
}
