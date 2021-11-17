package form3http

import (
	"encoding/json"
)

type Response struct {
	Status     string
	StatusCode int
	body       []byte
}

func (r *Response) Bytes() []byte {
	return r.body
}

func (r *Response) String() string {
	return string(r.body)
}

func (r *Response) UnmarshalJson(target interface{}) error {
	return json.Unmarshal(r.Bytes(), target)
}
