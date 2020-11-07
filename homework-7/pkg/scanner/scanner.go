// scanner - служба поискового робота
// Задача - сканировать страницу, находить на ней ссылки
package scanner

import (
	"fmt"
)

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

type Interface interface {
	Scan(url string, depth int, ch chan<- []Document)
}

type Service struct {
	Interface
}

func New(sc Interface) *Service {
	return &Service{sc}
}
