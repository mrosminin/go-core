// btree - реализация двоичного дерева поиска для хранения документов
package btree

import (
	"encoding/json"
	"errors"
	"go-core-own/homework-16/pkg/scanner"
	"sync"
)

// Tree - Двоичное дерево поиска
type Tree struct {
	root   *Element
	nextID int

	mux *sync.Mutex
}

// Element - элемент дерева
type Element struct {
	left, right *Element
	Doc         scanner.Document
}

func New() *Tree {
	return &Tree{
		mux: &sync.Mutex{},
	}
}

// Insert вставляет элемент в дерево
func (t *Tree) Insert(doc scanner.Document) (id int) {
	t.mux.Lock()
	defer t.mux.Unlock()

	if id = t.exists(doc); id != -1 {
		return id
	}
	doc.ID = t.nextID
	t.nextID = t.nextID + 1
	e := &Element{Doc: doc}
	if t.root == nil {
		t.root = e
		return doc.ID
	}
	insert(t.root, e)
	return doc.ID
}

// inset рекурсивно вставляет элемент в нужный уровень дерева
func insert(node *Element, new *Element) {
	if new.Doc.ID < node.Doc.ID {
		if node.left == nil {
			node.left = new
			return
		}
		insert(node.left, new)
	}
	if new.Doc.ID >= node.Doc.ID {
		if node.right == nil {
			node.right = new
			return
		}
		insert(node.right, new)
	}
}

// Find находит элемент в дереве по ID документа, возвращает документ
func (t *Tree) Find(id int) (scanner.Document, error) {
	el := find(t.root, id)
	if el == nil {
		return scanner.Document{}, errors.New("документ не найден")
	}
	return el.Doc, nil
}

func find(el *Element, id int) *Element {
	if el == nil {
		return nil
	}
	if el.Doc.ID == id {
		return el
	}
	if el.Doc.ID < id {
		return find(el.right, id)
	}
	return find(el.left, id)
}

// Serialize сериализует дерево в json строку
func (t *Tree) Json() (jsonData []byte, err error) {
	t.mux.Lock()
	defer t.mux.Unlock()

	var docs []scanner.Document
	for i := 0; i < t.nextID; i++ {
		doc, err := t.Find(i)
		if err != nil {
			continue
		}
		docs = append(docs, doc)
	}
	jsonData, err = json.Marshal(docs)
	if err != nil {
		return []byte{}, err
	}
	return jsonData, nil
}

// Проверка наличия документа в дереве, если есть возвращает id документа
func (t *Tree) exists(newDoc scanner.Document) (id int) {
	for i := 0; i < t.nextID; i++ {
		doc, _ := t.Find(i)
		if doc.Title == newDoc.Title || doc.URL == newDoc.URL {
			return doc.ID
		}
	}
	return -1
}
