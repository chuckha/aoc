package advent

import (
	"context"

	"github.com/chuckha/aoc/internal/core/domain"
)

type questionCache interface {
	Insert(*domain.Question) error
	Get(year, day int) (*domain.Question, error)
}

type questionInterface interface {
	Get(ctx context.Context, year, day int) (*domain.Question, error)
	Submit(ctx context.Context, year, day int, answer *domain.Answer) ([]byte, error)
}

type Service struct {
	QuestionCache     questionCache
	QuestionInterface questionInterface
}

func NewService(qc questionCache, qi questionInterface) *Service {
	return &Service{
		QuestionCache:     qc,
		QuestionInterface: qi,
	}
}

type GetDescriptionInput struct {
	Year, Day int
	Force     bool
}

func (s *Service) GetDescription(ctx context.Context, in *GetDescriptionInput) ([]byte, error) {
	q, err := s.getQuestion(ctx, in.Year, in.Day, in.Force)
	if err != nil {
		return nil, err
	}
	return q.RawQuestion, nil
}

func (s *Service) getQuestion(ctx context.Context, year, day int, force bool) (*domain.Question, error) {
	if force {
		q, err := s.QuestionInterface.Get(ctx, year, day)
		if err != nil {
			return nil, err
		}
		if err := s.QuestionCache.Insert(q); err != nil {
			return nil, err
		}
		return q, nil
	}

	q, err := s.QuestionCache.Get(year, day)
	if err != nil {
		return nil, err
	}
	if q == nil {
		q, err = s.QuestionInterface.Get(ctx, year, day)
		if err != nil {
			return nil, err
		}
		if err := s.QuestionCache.Insert(q); err != nil {
			return nil, err
		}
	}
	return q, nil
}

func (s *Service) GetInput(ctx context.Context, year, day int) ([]byte, error) {
	q, err := s.getQuestion(ctx, year, day, false)
	if err != nil {
		return nil, err
	}
	return q.Input.Input, nil
}

func (s *Service) SubmitAnswer(ctx context.Context, year, day, level int, answer string) ([]byte, error) {
	aswr, err := domain.NewAnswer(level, answer)
	if err != nil {
		return nil, err
	}
	return s.QuestionInterface.Submit(ctx, year, day, aswr)
}
