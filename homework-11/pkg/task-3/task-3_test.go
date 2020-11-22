package task_3

import (
	"fmt"
	"testing"
)

type mem struct {
	value string
}

func (m *mem) Write(p []byte) (n int, err error) {
	m.value += string(p)
	fmt.Println(m.value)
	return len(p), nil
}

func Test_writeStrings(t *testing.T) {
	var m mem
	args := []interface{}{1, 2, 3, "string4", 5, "string6"}
	writeStrings(&m, args...)
	want := "string4string6"

	if got := m.value; got != want {
		t.Errorf("writeStrings() = %v, want %v", got, want)
	}

}
