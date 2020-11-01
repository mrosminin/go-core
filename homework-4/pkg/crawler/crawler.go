package crawler

import "fmt"

// Scanner - интерфейс поискового робота.
type Scanner interface {
	Scan(url string, depth int) ([]Document, error)
}

// Document - документ, веб-страница, полученная поисковым роботом.
type Document struct {
	ID    int
	URL   string
	Title string
	Body  string
}

func (d Document) String() string {
	return fmt.Sprintf("%s: %s", d.URL, d.Title)
}
