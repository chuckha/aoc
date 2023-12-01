package web

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/chuckha/aoc/internal/core/domain"
	"gitlab.com/tozd/go/errors"
)

type serializer interface {
	Serialize(in interface{}) ([]byte, error)
}

type responseInterpreter interface {
	CorrectAnswer(data []byte) bool
}

type Advent struct {
	Token               domain.SessionToken
	Serializer          serializer
	ResponseInterpreter responseInterpreter
}

func NewAdvent(token string, serializer serializer, ri responseInterpreter) (*Advent, error) {
	tkn, err := domain.NewSessionToken(token)
	if err != nil {
		return nil, err
	}
	return &Advent{
		Token:               tkn,
		Serializer:          serializer,
		ResponseInterpreter: ri,
	}, nil
}

// TODO: use this line to figure out if the user is not logged in
// Puzzle inputs differ by user.  Please log in to get your puzzle input.

func (a *Advent) Get(ctx context.Context, year, day int) (*domain.Question, error) {
	link, err := url.Parse(fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	question, err := a.fetch(ctx, link)
	if err != nil {
		return nil, err
	}
	start := bytes.Index(question, []byte("<main>"))
	end := bytes.Index(question, []byte("</main>"))
	question = question[start : end+len("</main>")]
	link, err = url.Parse(fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	input, err := a.fetch(ctx, link)
	if err != nil {
		return nil, err
	}
	// examples are nil for now
	in, err := domain.NewInput(input, nil)
	if err != nil {
		return nil, err
	}
	return domain.NewQuestion(year, day, question, in)
}

func (a *Advent) Submit(ctx context.Context, year, day int, answer *domain.Answer) ([]byte, error) {
	link, err := url.Parse(fmt.Sprintf("https://adventofcode.com/%d/day/%d/answer", year, day))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	req, err := a.newReq(http.MethodPost, link.String())
	if err != nil {
		return nil, err
	}
	form := url.Values{
		"level":  []string{fmt.Sprintf("%d", answer.Level)},
		"answer": []string{answer.Answer},
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Body = io.NopCloser(strings.NewReader(form.Encode()))
	resp, err := doRequest(req)
	if err != nil {
		return nil, err
	}
	start := bytes.Index(resp, []byte("<main>"))
	end := bytes.Index(resp, []byte("</main>"))
	resp = resp[start : end+len("</main>")]

	if a.ResponseInterpreter.CorrectAnswer(resp) {
		return nil, nil
	}
	return resp, nil
}

func (a *Advent) fetch(ctx context.Context, link *url.URL) ([]byte, error) {
	req, err := a.newReq(http.MethodGet, link.String())
	if err != nil {
		return nil, err
	}
	return doRequest(req)
}

func doRequest(req *http.Request) ([]byte, error) {
	fmt.Fprintln(os.Stderr, "ABOUT TO MAKE A NETWORK CALL")
	fmt.Fprintln(os.Stderr, req.Method, req.URL.String())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	body, err := io.ReadAll(resp.Body)
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
	fmt.Println("debug", req)
	return req, nil
}
