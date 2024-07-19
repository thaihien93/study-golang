package routes

import (
	"todo_app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	taskGroup := r.Group("/tasks")
	{
		taskGroup.GET("", controllers.GetTasks)
		taskGroup.POST("/create", controllers.CreateTask)
		// taskGroup.PUT("/:id", controllers.UpdateTask)
		// taskGroup.DELETE("/:id", controllers.DeleteTask)
	}
}
