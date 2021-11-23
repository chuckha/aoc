package domain

import (
	"time"

	"github.com/pkg/errors"
)

type SessionToken string

func NewSessionToken(token string) (SessionToken, error) {
	if token == "" {
		return "", errors.New("token cannot be empty")
	}
	return SessionToken(token), nil
}

type Example struct {
	Input  []byte
	Answer string
}

func NewExample(input []byte, answer string) (*Example, error) {
	if input == nil {
		return nil, errors.New("input cannot be nil")
	}
	return &Example{
		Input:  input,
		Answer: answer,
	}, nil
}

type Input struct {
	Level    int
	Examples []Example
	Input    []byte
}

func NewInput(level int, input []byte, examples []Example) (*Input, error) {
	if level != 1 && level != 2 {
		return nil, errors.New("only level 1 and 2 are valid")
	}
	return &Input{
		Level:    level,
		Examples: examples,
		Input:    input,
	}, nil
}

type Question struct {
	Year        int
	Day         int
	RawQuestion []byte
	Inputs      []*Input
}

func NewQuestion(year, day int, raw []byte, inputs []*Input) (*Question, error) {
	if year < 2015 || year > time.Now().Year() {
		return nil, errors.New("year is invalid")
	}
	if day == 0 {
		return nil, errors.New("day must be bigger than 0")
	}
	if day > 31 {
		return nil, errors.New("day must be smaller than 31")
	}
	return &Question{
		Year:        year,
		Day:         day,
		RawQuestion: raw,
		Inputs:      inputs,
	}, nil
}
