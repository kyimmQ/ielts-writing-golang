package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kyimmQ/ielts-writing-golang/global"
	"github.com/kyimmQ/ielts-writing-golang/internal/modules/prompt"
)

func InitPromptRoute(api *gin.RouterGroup) {
	promptRepo := prompt.NewPromptRepository(global.MongoDB)
	promptService := prompt.NewPromptService(promptRepo)
	promptHandler := prompt.NewPromptHandler(promptService)

	promptRoute := api.Group("/prompts")
	{
		addGroupRoutes(promptRoute, getPromptRoute(promptHandler))
	}
}

func getPromptRoute(h prompt.PromptHandlerI) []Route {
	return []Route{
		{
			Method:  "POST",
			Path:    "",
			Handler: h.CreatePrompt,
		},
		{
			Method:  "GET",
			Path:    "/random",
			Handler: h.GetRandomPrompt,
		},
	}
}
