// btree - реализация двоичного дерева поиска для хранения документов
package btree

import (
	"encoding/json"
	"errors"
	"go-core-own/homework-6/pkg/scanner"
)

// Tree - Двоичное дерево поиска
type Tree struct {
	root   *Element
	lastID int
}

// Element - элемент дерева
type Element struct {
	left, right *Element
	Doc         scanner.Document
}

// Insert - вставка элемента в дерево
func (t *Tree) Insert(doc scanner.Document) (id int) {
	if id = t.exists(doc); id != 0 {
		return id
	}
	t.lastID = t.lastID + 1
	doc.ID = t.lastID
	e := &Element{Doc: doc}
	if t.root == nil {
		t.root = e
		return t.lastID
	}
	insert(t.root, e)
	return t.lastID
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

// Find - поиск элемента в дереве по ID документа, возвращает документ
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

// Сериализация дерева
func (t *Tree) Serialize() (jsonData []byte, err error) {
	var docs []scanner.Document
	for i := 0; i < t.lastID; i++ {
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
	for i := 0; i < t.lastID; i++ {
		doc, _ := t.Find(i)
		if doc.Title == newDoc.Title && doc.URL == newDoc.URL {
			return doc.ID
		}
	}
	return 0
}
