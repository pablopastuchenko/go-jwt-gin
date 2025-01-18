package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pablopastuchenko/go-jwt-gin/middleware"
	controller "github.com/pablopastuchenko/go-jwt-gin/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
