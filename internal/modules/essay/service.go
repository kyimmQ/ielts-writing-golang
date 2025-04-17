package essay

import (
	"context"

	"github.com/google/uuid"
	"github.com/kyimmQ/ielts-writing-golang/internal/entity"
	"github.com/kyimmQ/ielts-writing-golang/internal/modules/essay/dto"
	"github.com/kyimmQ/ielts-writing-golang/internal/modules/prompt"
)

type EssayServiceI interface {
	SubmitEssay(ctx context.Context, req *dto.CreateEssayRequest) error
	SaveDraft(ctx context.Context, req *dto.UpdateEssayRequest) error
	GetUserHistory(ctx context.Context) ([]dto.EssayWithPromptResponse, error)
	GetUserDrafts(ctx context.Context) ([]dto.EssayWithPromptResponse, error)
}

type EssayService struct {
	repo       EssayRepositoryI
	promptRepo prompt.PromptRepositoryI
}

func NewEssayService(repo EssayRepositoryI, promptRepo prompt.PromptRepositoryI) EssayServiceI {
	return &EssayService{repo: repo, promptRepo: promptRepo}
}

func (s *EssayService) SubmitEssay(ctx context.Context, req *dto.CreateEssayRequest) error {
	essay := req.ToEntity(ctx)
	return s.repo.CreateEssay(ctx, essay)
}

func (s *EssayService) SaveDraft(ctx context.Context, req *dto.UpdateEssayRequest) error {
	return s.repo.UpdateEssayDraft(ctx, req)
}

func (s *EssayService) GetUserHistory(ctx context.Context) ([]dto.EssayWithPromptResponse, error) {
	essays, err := s.repo.GetUserEssays(ctx, ctx.Value("userId").(uuid.UUID), []entity.EssayStatus{entity.StatusGraded, entity.StatusGrading})
	if err != nil {
		return nil, err
	}

	var essayResponses []dto.EssayWithPromptResponse
	for _, essay := range essays {
		prompt, err := s.promptRepo.GetPromptByID(ctx, essay.PromptID)
		if err != nil {
			return nil, err
		}
		essayResponse := dto.EssayWithPromptResponse{
			ID:        essay.ID,
			PromptID:  essay.PromptID,
			Prompt:    prompt.Prompt,
			Content:   essay.Content,
			Status:    string(essay.Status),
			Band:      essay.Band,
			TimeTaken: essay.TimeTaken,
			UpdatedAt: essay.UpdatedAt,
		}
		essayResponses = append(essayResponses, essayResponse)
	}
	return essayResponses, nil
}

func (s *EssayService) GetUserDrafts(ctx context.Context) ([]dto.EssayWithPromptResponse, error) {
	essays, err := s.repo.GetUserEssays(ctx, ctx.Value("userId").(uuid.UUID), []entity.EssayStatus{entity.StatusDraft})
	if err != nil {
		return nil, err
	}

	var essayResponses []dto.EssayWithPromptResponse
	for _, essay := range essays {
		prompt, err := s.promptRepo.GetPromptByID(ctx, essay.PromptID)
		if err != nil {
			return nil, err
		}
		essayResponse := dto.EssayWithPromptResponse{
			ID:        essay.ID,
			PromptID:  essay.PromptID,
			Prompt:    prompt.Prompt,
			Content:   essay.Content,
			Status:    string(essay.Status),
			Band:      essay.Band,
			TimeTaken: essay.TimeTaken,
			UpdatedAt: essay.UpdatedAt,
		}
		essayResponses = append(essayResponses, essayResponse)
	}
	return essayResponses, nil
}
