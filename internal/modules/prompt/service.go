package prompt

import (
	"context"

	"github.com/kyimmQ/ielts-writing-golang/internal/modules/prompt/dto"
)

type PromptServiceI interface {
	CreatePrompt(ctx context.Context, req *dto.CreatePromptRequest) error
	GetRandomPrompt(ctx context.Context) (*dto.PromptResponse, error)
}

type PromptService struct {
	promptRepo PromptRepositoryI
}

func NewPromptService(promptRepo PromptRepositoryI) PromptServiceI {
	return &PromptService{
		promptRepo: promptRepo,
	}
}

func (s *PromptService) CreatePrompt(ctx context.Context, req *dto.CreatePromptRequest) error {
	prompt := req.ToEntity()
	return s.promptRepo.CreatePrompt(ctx, prompt)
}

func (s *PromptService) GetRandomPrompt(ctx context.Context) (*dto.PromptResponse, error) {
	prompt, err := s.promptRepo.GetRandomPrompt(ctx)
	if err != nil {
		return nil, err
	}
	return &dto.PromptResponse{
		ID:     prompt.ID,
		Prompt: prompt.Prompt,
	}, nil
}
