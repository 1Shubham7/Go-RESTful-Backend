package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/1shubham7/jwt/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controller.Signup())
	incomingRoutes.POST("user/login", controllr.Login())
}