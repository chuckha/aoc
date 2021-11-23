package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/chuckha/aoc/internal/core/services/advent"
	"github.com/chuckha/aoc/internal/infrastructure/diskcache"
	"github.com/chuckha/aoc/internal/infrastructure/serde/json"
	"github.com/chuckha/aoc/internal/infrastructure/web"
)

func main() {
	ctx := context.Background()
	home := os.Getenv("HOME")
	dir := filepath.Join(home, ".aoc")
	tkn, err := os.ReadFile(dir + "/token")
	if err != nil {
		panic(err)
	}

	serde := json.NewSerDe()
	qc := diskcache.NewQuestionsCache(dir, serde)
	qf, err := web.NewAdvent(string(tkn))
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}
	svc := advent.NewService(qc, qf)
	input, err := svc.GetInput(ctx, 2015, 1, 1)
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}
	fmt.Println(string(input))
}
