package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/1shubham7/jwt/controllers"
)

// this is when the user has not signed up. userRouter is when the user has logged in
// and has the token.
func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.Signup())
	incomingRoutes.POST("user/login", controllers.Login())
}