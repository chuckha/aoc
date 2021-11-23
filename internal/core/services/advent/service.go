package advent

import (
	"context"
	"errors"

	"github.com/chuckha/aoc/internal/core/domain"
)

type questionCache interface {
	Insert(*domain.Question) error
	Get(year, day int) (*domain.Question, error)
}

type questionFetcher interface {
	Get(ctx context.Context, year, day int) (*domain.Question, error)
}

type Service struct {
	QuestionCache   questionCache
	QuestionFetcher questionFetcher
}

func NewService(qc questionCache, qf questionFetcher) *Service {
	return &Service{
		QuestionCache:   qc,
		QuestionFetcher: qf,
	}
}

func (s *Service) GetInput(ctx context.Context, year, day, level int) ([]byte, error) {
	q, err := s.QuestionCache.Get(year, day)
	if err != nil {
		return nil, err
	}
	if q == nil {
		q, err = s.QuestionFetcher.Get(ctx, year, day)
		if err != nil {
			return nil, err
		}
		if err := s.QuestionCache.Insert(q); err != nil {
			return nil, err
		}
	}
	for _, input := range q.Inputs {
		if input.Level == level {
			return input.Input, nil
		}
	}
	return nil, errors.New("input level not found")
}
