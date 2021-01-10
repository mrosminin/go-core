// spider реализует сканер содержимого веб-сайтов.
// Пакет позволяет получить список ссылок и заголовков страниц внутри веб-сайта по его URL.
package spider

import (
	"go-core-own/homework-19/pkg/scanner"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

type Service struct {
	Interface
}

func New(n Interface) *Service {
	return &Service{n}
}

// Scan осуществляет рекурсивный обход ссылок сайта, указанного в URL,
// с учётом глубины перехода по ссылкам, переданной в depth.
func (s *Service) Scan(url string, depth int) (data []scanner.Document, err error) {
	pages := make(map[string]string)
	err = parse(s.Interface, url, url, depth, pages)
	if err != nil {
		return []scanner.Document{}, err
	}

	for url, title := range pages {
		item := scanner.Document{
			URL:   url,
			Title: title,
		}
		data = append(data, item)
	}
	return data, nil
}

// parse рекурсивно обходит ссылки на странице, переданной в url.
// Глубина рекурсии задаётся в depth.
// Каждая найденная ссылка записывается в ассоциативный массив
// data вместе с названием страницы.
func parse(n Interface, url, baseurl string, depth int, data map[string]string) error {
	if depth == 0 {
		return nil
	}

	page, err := n.Page(url)
	if err != nil {
		return err
	}

	data[url] = pageTitle(page)

	links := pageLinks(nil, page)
	for _, link := range links {
		// ссылка уже отсканирована
		if data[link] != "" {
			continue
		}
		// ссылка содержит базовый url полностью
		if strings.HasPrefix(link, baseurl) {
			parse(n, link, baseurl, depth-1, data)
		}
		// относительная ссылка
		if strings.HasPrefix(link, "/") && len(link) > 1 {
			next := baseurl + link
			parse(n, next, baseurl, depth-1, data)
		}
	}

	return nil
}

// pageTitle осуществляет рекурсивный обход HTML-страницы и возвращает значение элемента <tittle>.
func pageTitle(n *html.Node) string {
	var title string
	if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
		return n.FirstChild.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		title = pageTitle(c)
		if title != "" {
			break
		}
	}
	return title
}

// pageLinks рекурсивно сканирует узлы HTML-страницы и возвращает все найденные ссылки без дубликатов.
func pageLinks(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				if !contains(links, a.Val) {
					links = append(links, a.Val)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = pageLinks(links, c)
	}
	return links
}

// contains возвращает true если массив содержит переданное значение
func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// Интерфейс для подмены зависимости от сети
type Interface interface {
	Page(url string) (*html.Node, error)
}

type Net struct{}

// page получает страницу по ссылке и раскодирует ее
func (n *Net) Page(url string) (*html.Node, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	page, err := html.Parse(response.Body)
	if err != nil {
		return nil, err
	}

	return page, nil
}
