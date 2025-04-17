package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kyimmQ/ielts-writing-golang/global"
	"github.com/kyimmQ/ielts-writing-golang/internal/modules/essay"
	"github.com/kyimmQ/ielts-writing-golang/internal/modules/prompt"
)

func InitEssayRoute(api *gin.RouterGroup) {
	promptRepo := prompt.NewPromptRepository(global.MongoDB)
	essayRepo := essay.NewEssayRepository(global.MongoDB)
	essayService := essay.NewEssayService(essayRepo, promptRepo)
	essayHandler := essay.NewEssayHandler(essayService)

	essayRoute := api.Group("/essay")
	{
		addGroupRoutes(essayRoute, getEssayRoutes(essayHandler))
	}
}

func getEssayRoutes(h essay.EssayHandlerI) []Route {
	return []Route{
		{
			Method:  "POST",
			Path:    "/submit",
			Handler: h.SubmitEssay,
		},
		{
			Method:  "PUT",
			Path:    "/draft",
			Handler: h.SaveDraft,
		},
		{
			Method:  "GET",
			Path:    "/history",
			Handler: h.GetUserHistory,
		},
		{
			Method:  "GET",
			Path:    "/drafts",
			Handler: h.GetUserDrafts,
		},
	}
}
