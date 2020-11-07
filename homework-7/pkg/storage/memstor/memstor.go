// memstor - служба сохранения результатов сканирования в памяти (заглушка для тестов)
package memstor

type Memstor struct {
	memory []byte
}

func New() *Memstor {
	return &Memstor{}
}

func (ms *Memstor) Save(p []byte) error {
	ms.memory = p
	return nil
}

// Load - читает из переменной
func (ms *Memstor) Load() (p []byte, err error) {
	return ms.memory, nil
}
