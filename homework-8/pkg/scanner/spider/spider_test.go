package spider

import (
	"go-core-own/homework-8/pkg/scanner"
	"golang.org/x/net/html"
	"reflect"
	"testing"
)

func Test_pageTitle(t *testing.T) {
	want := "Заголовок"
	n := &html.Node{
		FirstChild: &html.Node{
			Type: html.ElementNode,
			Data: "title",
			FirstChild: &html.Node{
				Data: want,
			},
		},
	}
	if got := pageTitle(n); got != want {
		t.Errorf("pageTitle() = %v, want %v", got, want)
	}
}

func Test_pageLinks(t *testing.T) {
	n := &html.Node{
		Type: html.ElementNode,
		Data: "a",
		Attr: []html.Attribute{
			{
				Key: "href",
				Val: "www.ya.ru",
			},
			{
				Key: "href",
				Val: "www.yandex.ru",
			},
			{
				Key: "href",
				Val: "www.ya.ru",
			},
		},
		FirstChild: &html.Node{
			Type: html.ElementNode,
			Data: "a",
			Attr: []html.Attribute{
				{
					Key: "href",
					Val: "www.ya.ru",
				},
				{
					Key: "href",
					Val: "www.transflow.ru",
				},
			},
		},
	}
	want := []string{
		"www.ya.ru",
		"www.yandex.ru",
		"www.transflow.ru",
	}
	if got := pageLinks([]string{}, n); !reflect.DeepEqual(got, want) {
		t.Errorf("pageTitle() = %v, want %v", got, want)
	}
}

type fake struct{}

func (*fake) Page(string) (*html.Node, error) {
	n := &html.Node{
		Type: html.ElementNode,
		Data: "title",
		FirstChild: &html.Node{
			Data: "Заголовок",
			NextSibling: &html.Node{
				Type: html.ElementNode,
				Data: "a",
				Attr: []html.Attribute{
					{
						Key: "href",
						Val: "www.site.ru/link1",
					},
					{
						Key: "href",
						Val: "/link2",
					},
					{
						Key: "href",
						Val: "/link1",
					},
					{
						Key: "href",
						Val: "www.site.ru/link3",
					},
					{
						Key: "href",
						Val: "/link3",
					},
				},
			},
		},
	}
	return n, nil
}

func Test_parse(t *testing.T) {
	got := make(map[string]string)
	_ = parse(&fake{}, "www.site.ru", "www.site.ru", 2, got)
	want := map[string]string{
		"www.site.ru":       "Заголовок",
		"www.site.ru/link1": "Заголовок",
		"www.site.ru/link2": "Заголовок",
		"www.site.ru/link3": "Заголовок",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("parse(). pages = %v, want %v", got, want)
	}
}

func TestService_Scan(t *testing.T) {
	s := Service{&fake{}}
	got, _ := s.Scan("www.site.ru", 2)
	want := []scanner.Document{
		scanner.Document{ID: 0, URL: "www.site.ru", Title: "Заголовок", Body: ""},
		scanner.Document{ID: 0, URL: "www.site.ru/link1", Title: "Заголовок", Body: ""},
		scanner.Document{ID: 0, URL: "www.site.ru/link2", Title: "Заголовок", Body: ""},
		scanner.Document{ID: 0, URL: "www.site.ru/link3", Title: "Заголовок", Body: ""},
	}
	for _, wd := range want {
		found := false
		for _, gd := range got {
			if reflect.DeepEqual(wd, gd) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Scan() = %v, want %v", got, want)
			break
		}
	}
}
