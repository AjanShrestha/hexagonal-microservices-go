package shortener

// Redirect is the main model for URL Shortener
//
// Code: string Key to fetch the URL
// URL: string The actual URL
// CreatedAt: int64 Timestamp when the Code is created
type Redirect struct {
	Code      string
	URL       string
	CreatedAt int64
}
