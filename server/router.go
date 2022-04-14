package server

import (
	"gin-todo-app/controllers"
	"gin-todo-app/middlewares"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
)

func InitializeRouter(app firebase.App, client firestore.Client) *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.FirebaseMiddleware(app, client))

	todoGroup := router.Group("todo")
	{
		todoGroup.POST("create", controllers.CreateToDo)
		todoGroup.GET("getall", controllers.GetAllToDo)
		todoGroup.GET("get/:id", controllers.GetToDo)
		todoGroup.PUT("update/:id", controllers.UpdateToDo)
		todoGroup.DELETE("delete/:id", controllers.DeleteToDo)
	}

	return router
}
