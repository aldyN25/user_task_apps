package apiv1

import (
	"context"
	"log"
	"net/http"
	tasksController "user_task_apps/app/controllers/tasks"
	usersController "user_task_apps/app/controllers/users"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig *oauth2.Config
	// oauthStateString  = "random"
)

func init() {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/callback",
		ClientID:     "240461886825-2bfadd3if19jo58qu3j77suuaguc70rb.apps.googleusercontent.com",
		ClientSecret: "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",
		Scopes:       []string{"openid", "profile"},
		Endpoint:     google.Endpoint,
	}
}

func ApiRoutes(router *gin.Engine) {

	r := router.Group("/api/v1")

	usersGroup := r.Group("/users")
	tasksGroup := r.Group("/tasks")

	usersGroup.POST("/", usersController.CreateUsers)
	usersGroup.GET("/", usersController.GetAllUsers)
	usersGroup.GET("/:id", usersController.GetUsersById)
	usersGroup.PUT("/:id", usersController.UpdateUsers)
	usersGroup.DELETE("/:id", usersController.DeleteUsers)

	router.GET("/auth/login", func(c *gin.Context) {
		url := googleOauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
		log.Println(http.StatusTemporaryRedirect, url)

		c.Redirect(http.StatusTemporaryRedirect, url)
	})

	router.GET("/auth/callback", func(c *gin.Context) {
		code := c.Query("code")
		token, err := googleOauthConfig.Exchange(context.Background(), code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"access_token": token.AccessToken, "token_type": token.TokenType})
	})

	// Middleware for OAuth2 authentication
	tasksGroup.Use(func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		_, err := googleOauthConfig.TokenSource(context.Background(), &oauth2.Token{AccessToken: tokenString}).Token()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	})

	tasksGroup.POST("/", tasksController.CreateTasks)
	tasksGroup.GET("/", tasksController.GetAllTasks)
	tasksGroup.GET("/:id", tasksController.GetTasksById)
	tasksGroup.PUT("/:id", tasksController.UpdateTasks)
	tasksGroup.DELETE("/:id", tasksController.DeleteTasks)
}
