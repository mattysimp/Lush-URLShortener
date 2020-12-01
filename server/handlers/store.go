package handlers

// URLStore interface of URL
type URLStore interface {
	SetURL(*URL) (err error)
	GetURL(string) (*URL, error)
}
