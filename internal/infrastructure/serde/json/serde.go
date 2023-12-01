package json

import (
	"encoding/json"

	"gitlab.com/tozd/go/errors"
)

type SerDe struct{}

func NewSerDe() *SerDe {
	return &SerDe{}
}

func (s *SerDe) Serialize(in interface{}) ([]byte, error) {
	b, err := json.Marshal(in)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return b, nil
}

func (s *SerDe) Deserialize(data []byte, in interface{}) error {
	if err := json.Unmarshal(data, in); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
