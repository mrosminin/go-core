// diskstor - служба хранения данных на диске
package diskstor

import (
	"io/ioutil"
	"os"
	"sync"
)

type Diskstor struct {
	file *os.File
	mux  sync.Mutex
}

// New - конструктор службы, создает файл для хранения данных
func New(filename string) (*Diskstor, error) {
	var f *os.File
	f, err := os.Open(filename)
	if err != nil {
		f, err = os.Create(filename)
		if err != nil {
			return nil, err
		}
	}
	return &Diskstor{file: f}, nil
}

// Save - пишет строку в файл
func (ds *Diskstor) Save(p []byte) error {
	ds.mux.Lock()
	defer ds.mux.Unlock()
	err := ioutil.WriteFile(ds.file.Name(), p, 0666)
	if err != nil {
		return err
	}
	return nil
}

// Load - читает строку из файла
func (ds *Diskstor) Load() (p []byte, err error) {
	ds.mux.Lock()
	defer ds.mux.Unlock()
	f, err := os.Open(ds.file.Name())
	defer f.Close()
	p, err = ioutil.ReadFile(f.Name())
	if err != nil {
		return nil, err
	}
	return p, nil
}
