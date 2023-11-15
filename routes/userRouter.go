package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/1shubham7/jwt/controllers"
	middleware "github.com/1shubham7/jwt/middleware"
)

// user should not be able to use userRoute without the token
func UserRoutes (incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenthicate())
	// user routes are public routes but these must be authenticated, that
	// is why we have Authenticate() before these
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("users/:user_id", controllers.GetUserById())
}