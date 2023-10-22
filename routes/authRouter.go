package routes

import (
	controller "jwt-authentication/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controller.signup())
	incomingRoutes.POST("users/login", controller.Login())
}
