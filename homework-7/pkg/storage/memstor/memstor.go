package memstor

// MStorage - служба сохранения результатов сканирования в памяти (заглушка для тестов)
type Memstor struct {
	memory []byte
}

// New - конструктор службы
func New() *Memstor {
	return &Memstor{}
}

// Save - пишет в переменную
func (ms *Memstor) Save(p []byte) error {
	ms.memory = p
	return nil
}

// Load - читает из переменной
func (ms *Memstor) Load() (p []byte, err error) {
	return ms.memory, nil
}
