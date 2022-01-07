package local

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

type Adapter struct {
	homeDir string
}

func NewAdapter(homeDir string) *Adapter {
	return &Adapter{homeDir: homeDir}
}

func (a *Adapter) GetToken() ([]byte, error) {
	dir := filepath.Join(a.homeDir, ".aoc")
	tkn, err := os.ReadFile(dir + "/token")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return tkn, nil
}
