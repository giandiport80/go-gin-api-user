package main

import (
	"log"
	"net/http"
	"os"

	"github.com/giandiport80/go-gin-api-user/config"
	"github.com/giandiport80/go-gin-api-user/models"
	"github.com/giandiport80/go-gin-api-user/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{})

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Hello world",
		})
	})

	routes.UserRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default fallback
	}

	log.Printf("Server running at http://localhost:%s", port)
	r.Run(":" + port)
}
