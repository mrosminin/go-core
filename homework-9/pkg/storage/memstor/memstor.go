// memstor - служба сохранения результатов сканирования в памяти (заглушка для тестов)
package memstor

import (
	"errors"
)

type Memstor struct {
	memory []byte
	err    bool
}

func New(err bool) *Memstor {
	return &Memstor{err: err}
}

func (ms *Memstor) Save(p []byte) error {
	if ms.err {
		return errors.New("тестовая ошибка memstor.Save()")
	}
	ms.memory = p
	return nil
}

// Load - читает из переменной
func (ms *Memstor) Load() (p []byte, err error) {
	if ms.err {
		return []byte{}, errors.New("тестовая ошибка memstor.Load()")
	}
	return ms.memory, nil
}
