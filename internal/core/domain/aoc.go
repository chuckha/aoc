package domain

import (
	"time"

	"github.com/pkg/errors"
)

type Answer struct {
	Level  int    `json:"level"`
	Answer string `json:"answer"`
}

func NewAnswer(level int, answer string) (*Answer, error) {
	if level > 2 {
		return nil, errors.New("undefined level greater than 2")
	}
	if answer == "" {
		return nil, errors.New("very unlikely the answer is an empty string")
	}
	return &Answer{
		Level:  level,
		Answer: answer,
	}, nil
}

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
	Examples []Example
	Input    []byte
}

func NewInput(input []byte, examples []Example) (*Input, error) {
	return &Input{
		Examples: examples,
		Input:    input,
	}, nil
}

type Question struct {
	Year        int
	Day         int
	RawQuestion []byte
	Input       *Input
}

func NewQuestion(year, day int, raw []byte, input *Input) (*Question, error) {
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
		Input:       input,
	}, nil
}
