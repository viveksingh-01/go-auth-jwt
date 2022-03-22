package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/viveksingh-01/go-auth-jwt/controllers"
)

func AuthRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("user/singup", controller.Signup())
	incomingRoutes.POST("user/login", controller.Login())
}