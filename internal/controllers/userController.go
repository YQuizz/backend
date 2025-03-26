package controllers

import (
	"net/http"
	"os"
	"time"
	"yquiz_back/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(c *gin.Context, user *models.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid":    user.ID,
		"firstname": user.FirstName,
		"lastname":  user.LastName,
		"email":     user.Email,
		"role":      user.Role,
		"class_id":  user.ClassID,
		"exp":       time.Now().Add(time.Hour * 24 * 10).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_JWT")))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create token"})
		return ""
	}
	// send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*10, "", "", true, true)
	return tokenString
}
