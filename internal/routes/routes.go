package routes

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/kyimmQ/ielts-writing-golang/docs"
	middleware "github.com/kyimmQ/ielts-writing-golang/internal/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

func InitRoute() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Authenticated API routes
	api := r.Group("/api")
	{
		InitAuthRoute(api)
	}
	authenApi := api.Group("")
	authenApi.Use(middleware.AuthenMiddleware())
	{
		InitPromptRoute(authenApi)
		InitEssayRoute(authenApi)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

func addGroupRoutes(g *gin.RouterGroup, routes []Route) {
	for _, route := range routes {
		switch route.Method {
		case "GET":
			g.GET(route.Path, route.Handler)
		case "POST":
			g.POST(route.Path, route.Handler)
		case "PUT":
			g.PUT(route.Path, route.Handler)
		case "DELETE":
			g.DELETE(route.Path, route.Handler)
		}
	}
}
