package json

import (
	"encoding/json"

	"github.com/AjanShrestha/hexagonal-microservices-go/shortener"
	"github.com/pkg/errors"
)

// Redirect struct that implements Serializer interface
type Redirect struct{}

// Decode takes input and generates a redirect
func (r *Redirect) Decode(input []byte) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	if err := json.Unmarshal(input, redirect); err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Decode")
	}
	return redirect, nil
}

// Encode takes input and returns the raw value
func (r *Redirect) Encode(input *shortener.Redirect) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Encode")
	}
	return rawMsg, nil
}
