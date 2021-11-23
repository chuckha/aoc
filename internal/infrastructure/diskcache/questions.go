package diskcache

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/chuckha/aoc/internal/core/domain"
	"github.com/pkg/errors"
)

type serde interface {
	Serialize(interface{}) ([]byte, error)
	Deserialize([]byte, interface{}) error
}

type QuestionsCache struct {
	Directory string
	SerDe     serde
}

func NewQuestionsCache(dir string, serde serde) *QuestionsCache {
	return &QuestionsCache{
		Directory: dir,
		SerDe:     serde,
	}
}

func (q *QuestionsCache) Insert(question *domain.Question) error {
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

func (q *QuestionsCache) Get(year, day int) (*domain.Question, error) {
	file := filepath.Join(q.Directory, toPath(year, day), "data")
	_, err := os.Stat(file)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}
		return nil, errors.WithStack(err)
	}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	question := &domain.Question{}
	if err := q.SerDe.Deserialize(data, question); err != nil {
		return nil, err
	}
	return question, nil
}

func toPath(year, day int) string {
	return filepath.Join(fmt.Sprintf("%d", year), fmt.Sprintf("%d", day))
}
