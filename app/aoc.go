package app

import (
	"os"
	"path/filepath"

	"github.com/chuckha/aoc/internal/adapters/local"
	"github.com/chuckha/aoc/internal/core/services/advent"
	"github.com/chuckha/aoc/internal/infrastructure/diskcache"
	"github.com/chuckha/aoc/internal/infrastructure/serde/json"
	"github.com/chuckha/aoc/internal/infrastructure/web"
)

// Application is a combination of infrastructure adapters & services to make
// a reasonably functioning that contains all the business logic.
// All that is left is to wrap the output/input with
// various adapters (cli adapter).
// This is the main entry point to programatically access the business logic.
type Application struct {
	*advent.Service
}

func NewApplication() (*Application, error) {
	home := os.Getenv("HOME")
	adapter := local.NewAdapter(home)
	serde := json.NewSerDe()
	dir := filepath.Join(home, ".aoc")
	tkn, err := adapter.GetToken()
	if err != nil {
		return nil, err
	}
	qc := diskcache.NewQuestionsCache(dir, serde)
	qi, err := web.NewAdvent(string(tkn), serde, &web.ResponseInterpreter{})
	if err != nil {
		return nil, err
	}
	svc := advent.NewService(qc, qi)
	return &Application{svc}, nil
}
