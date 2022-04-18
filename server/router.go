package server

import (
	"gin-todo-app/controllers"
	"gin-todo-app/middlewares"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           To Do API
// @version         2.0
// @description     API to practice gin and gcp
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /todo
func InitializeRouter(app firebase.App, client firestore.Client, redis redis.Pool) *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.SetMiddleware(app, client, redis))

	todoGroup := router.Group("todo")
	{
		todoGroup.POST("create", controllers.CreateToDo)
		todoGroup.GET("getall", controllers.GetAllToDo)
		todoGroup.GET("get/:id", controllers.GetToDo)
		todoGroup.PUT("update/:id", controllers.UpdateToDo)
		todoGroup.DELETE("delete/:id", controllers.DeleteToDo)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
