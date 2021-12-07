package web

import "bytes"

type ResponseInterpreter struct{}

func (r *ResponseInterpreter) CorrectAnswer(data []byte) bool {
	return bytes.Contains(data, []byte("That's the right answer!"))
}
