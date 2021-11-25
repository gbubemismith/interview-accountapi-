package form3http

import (
	"encoding/json"
)

type Response struct {
	Status     string
	StatusCode int
	Body       []byte
}

func (r *Response) Bytes() []byte {
	return r.Body
}

func (r *Response) UnmarshalJson(target interface{}) error {
	return json.Unmarshal(r.Bytes(), target)
}
