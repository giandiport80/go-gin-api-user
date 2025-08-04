package routes

import (
	"github.com/giandiport80/go-gin-api-user/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	user := r.Group("/api/users")
	{
		user.GET("/", controllers.GetUsers)
		user.GET("/:id", controllers.GetUserById)
		user.POST("/", controllers.CreateUser)
		user.PUT("/:id", controllers.UpdateUser)
		user.DELETE("/:id", controllers.DeleteUser)
	}
}
