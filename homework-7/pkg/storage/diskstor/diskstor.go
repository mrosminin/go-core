package diskstor

import (
	"io/ioutil"
	"os"
)

// DStorage - служба хранения данных на диске
type Diskstor struct {
	file *os.File
}

// New - конструктор службы, создает файл для хранения данных
func New() (*Diskstor, error) {
	filename := "./diskstor.txt"
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

// Save - пишет json в файл
func (ds *Diskstor) Save(p []byte) error {
	err := ioutil.WriteFile(ds.file.Name(), p, 0666)
	if err != nil {
		return err
	}
	return nil
}

// Load - читает json из файла
func (ds *Diskstor) Load() (p []byte, err error) {
	f, err := os.Open(ds.file.Name())
	defer f.Close()
	p, err = ioutil.ReadFile(f.Name())
	if err != nil {
		return nil, err
	}
	return p, nil
}
