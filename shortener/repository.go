package shortener

// RedirectRepository defines interface to store
// Same function signature as service
type RedirectRepository interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}
