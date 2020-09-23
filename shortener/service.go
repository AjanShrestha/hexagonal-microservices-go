package shortener

// RedirectService is the main service interface
//
// In order to fulfill the interface, it has to implement
// Find(code string) (*Redirect, error) : Return the URL from code
// Store(redirect *Redirect) error : Generate code and store it
type RedirectService interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}
