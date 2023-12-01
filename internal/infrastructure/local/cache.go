package local

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/chuckha/aoc/internal/core/domain"
	"gitlab.com/tozd/go/errors"
)

type serde interface {
	Serialize(interface{}) ([]byte, error)
	Deserialize([]byte, interface{}) error
}

type questionGetter interface {
	Get(ctx context.Context, year, day int) (*domain.Question, error)
}

type Cache struct {
	Directory string
	SerDe     serde
	qg        questionGetter
}

func NewCache(dir string, serde serde, qg questionGetter) *Cache {
	return &Cache{
		Directory: dir,
		SerDe:     serde,
		qg:        qg,
	}
}

func (q *Cache) insert(question *domain.Question) error {
	dir := filepath.Join(q.Directory, toPath(question.Year, question.Day))
	if err := os.MkdirAll(dir, os.ModeDir|0o700); err != nil {
		return errors.WithStack(err)
	}

	// save it to `~/.aoc/<year>/<day>/data}
	data, err := q.SerDe.Serialize(question)
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(dir, "data"), data, 0o700); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (q *Cache) Get(ctx context.Context, year, day int) (*domain.Question, error) {
	// look in the cache first, if it's a hit, return, if it's a miss go get it
	file := filepath.Join(q.Directory, toPath(year, day), "data")
	_, err := os.Stat(file)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			question, err := q.qg.Get(ctx, year, day)
			if err != nil {
				return nil, err
			}
			if err := q.insert(question); err != nil {
				return nil, err
			}
			return question, nil
		}
		return nil, errors.WithStack(err)
	}
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	question := &domain.Question{}
	if err := q.SerDe.Deserialize(data, question); err != nil {
		return nil, err
	}
	return question, nil
}

func (q *Cache) Delete(year, day int) error {
	dir := filepath.Join(q.Directory, toPath(year, day))
	if err := os.RemoveAll(dir); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func toPath(year, day int) string {
	return filepath.Join(fmt.Sprintf("%d", year), fmt.Sprintf("%d", day))
}
