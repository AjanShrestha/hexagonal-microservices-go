package shortener

// RedirectSerializer is the serializer interface
//
// Decode(input []byte) (*Redirect, error) : It will take a slice of byte and output a tuple of Redirect and error
// Encode(input *Redirect) ([]byte, error) : It will take Redirect and output a tuple of slice of byte and error
type RedirectSerializer interface {
	Decode(input []byte) (*Redirect, error)
	Encode(input *Redirect) ([]byte, error)
}
