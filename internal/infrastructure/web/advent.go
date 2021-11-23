package web

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/chuckha/aoc/internal/core/domain"
	"github.com/pkg/errors"
)

type Advent struct {
	Token domain.SessionToken
}

func NewAdvent(token string) (*Advent, error) {
	tkn, err := domain.NewSessionToken(token)
	if err != nil {
		return nil, err
	}
	return &Advent{
		Token: tkn,
	}, nil
}

func (a *Advent) Get(ctx context.Context, year, day int) (*domain.Question, error) {
	link, err := url.Parse(fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	question, err := a.fetch(ctx, link)
	if err != nil {
		return nil, err
	}
	link, err = url.Parse(fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	input, err := a.fetch(ctx, link)
	if err != nil {
		return nil, err
	}
	in, err := domain.NewInput(1, input, nil)
	if err != nil {
		return nil, err
	}
	return domain.NewQuestion(year, day, question, []*domain.Input{in})
}

func (a *Advent) SubmitAnswer() {}

func (a *Advent) fetch(ctx context.Context, link *url.URL) ([]byte, error) {
	fmt.Fprintln(os.Stderr, "AOBUT TO MAKE A NETWORK CALL")
	req, err := a.newReq(http.MethodGet, link.String())
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return body, nil
}

func (a *Advent) newReq(method string, link string) (*http.Request, error) {
	req, err := http.NewRequest(method, link, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: string(a.Token),
	})
	return req, nil
}
