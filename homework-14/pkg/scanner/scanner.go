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
