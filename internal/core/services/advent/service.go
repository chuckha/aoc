package advent

import (
	"context"

	"github.com/chuckha/aoc/internal/core/domain"
)

type cachedReader interface {
	Get(ctx context.Context, year, day int) (*domain.Question, error)
	Delete(year, day int) error
}

type writer interface {
	Submit(ctx context.Context, year, day int, answer *domain.Answer) ([]byte, error)
}

type Service struct {
	reader cachedReader
	writer writer
}

func NewService(r cachedReader, w writer) *Service {
	return &Service{
		reader: r,
		writer: w,
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
		s.reader.Delete(year, day)
	}
	return s.reader.Get(ctx, year, day)
}

func (s *Service) GetInput(ctx context.Context, year, day int) ([]byte, error) {
	q, err := s.getQuestion(ctx, year, day, false)
	if err != nil {
		return nil, err
	}
	return q.Input.Input, nil
}

func (s *Service) SubmitAnswer(ctx context.Context, year, day, level int, answer string) ([]byte, error) {
	ans, err := domain.NewAnswer(level, answer)
	if err != nil {
		return nil, err
	}
	return s.writer.Submit(ctx, year, day, ans)
}
