package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kyimmQ/ielts-writing-golang/global"
	"github.com/kyimmQ/ielts-writing-golang/internal/modules/auth"
	"github.com/kyimmQ/ielts-writing-golang/internal/modules/user"
)

func InitAuthRoute(api *gin.RouterGroup) {
	userRepo := user.NewUserRepository(global.MongoDB)
	userService := user.NewUserService(userRepo)
	authService := auth.NewAuthService(userService)
	authHandler := auth.NewAuthHandler(authService)

	authRoute := api.Group("/auth")
	{
		addGroupRoutes(authRoute, getAuthRoute(authHandler))
	}
}

func getAuthRoute(h auth.AuthHandlerI) []Route {
	return []Route{
		{
			Method:  "POST",
			Path:    "/signup",
			Handler: h.SignUp,
		},
		{
			Method:  "POST",
			Path:    "/signin",
			Handler: h.SignIn,
		},
	}
}
