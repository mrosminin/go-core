package task_3

import (
	"io"
)

// writeStringsge выводит только строки в канал вывода
func writeStrings(w io.Writer, args ...interface{}) {
	for _, arg := range args {
		if str, ok := arg.(string); ok {
			_, _ = w.Write([]byte(str))
		}
	}
}
