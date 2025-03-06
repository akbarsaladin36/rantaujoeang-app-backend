package middleware

import (
	"net/http"
	"rantaujoeang-app-backend/database"
	"rantaujoeang-app-backend/models"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		checkSession, errCheckSession := checkSessionToken(tokenString)

		if errCheckSession != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  "401",
				"message": "Token is not found!",
			})
			c.Abort()
			return
		}

		if time.Now().After(checkSession.SessionExpiredAt) {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  "401",
				"message": "Token is expired!",
			})
			c.Abort()
			return
		}

		currentUser := map[string]string{
			"user_uuid":     checkSession.SessionUserUUID,
			"user_username": checkSession.SessionUserUsername,
			"user_role":     checkSession.SessionUserRole,
		}

		c.Set("currentUser", currentUser)
		c.Next()
	}
}

func IsAdminAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		getSessionUser, _ := c.Get("currentUser")

		currentUser := getSessionUser.(map[string]string)

		user_role := currentUser["user_role"]

		if user_role != "admin" {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  "401",
				"message": "This url can be accessed by admin!",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func IsUserAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		getSessionUser, _ := c.Get("currentUser")

		currentUser := getSessionUser.(map[string]string)

		user_role := currentUser["user_role"]

		if user_role != "user" {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  "401",
				"message": "This url can be accessed by user!",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func CurrentUser(c *gin.Context) map[string]string {
	getSessionUser, _ := c.Get("currentUser")

	currentUser := getSessionUser.(map[string]string)

	return currentUser
}

func checkSessionToken(tokenString string) (models.Session, error) {
	var session models.Session

	err := database.DB.Where("session_token = ?", tokenString).First(&session).Error

	return session, err
}

func CheckSessionUser(user_uuid string) (models.Session, error) {
	var session models.Session

	err := database.DB.Where("session_user_uuid = ?", user_uuid).First(&session).Error

	return session, err
}

func CreateSessionUser(session models.Session) (models.Session, error) {
	err := database.DB.Create(&session).Error

	return session, err
}

func UpdateSessionUser(session models.Session) (models.Session, error) {
	err := database.DB.Save(&session).Error

	return session, err
}
