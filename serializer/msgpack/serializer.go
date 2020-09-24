package msgpack

import (
	"github.com/AjanShrestha/hexagonal-microservices-go/shortener"
	"github.com/pkg/errors"
	"github.com/vmihailenco/msgpack/v5"
)

// Redirect struct that implements Serializer interface
type Redirect struct{}

// Decode takes input, decodes and generates a redirect
func (r *Redirect) Decode(input []byte) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	if err := msgpack.Unmarshal(input, redirect); err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Decode")
	}
	return redirect, nil
}

// Encode takes input, encode and returns the raw value
func (r *Redirect) Encode(input *shortener.Redirect) ([]byte, error) {
	rawMsg, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Encode")
	}
	return rawMsg, nil
}
