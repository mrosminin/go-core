package task_2

import (
	"reflect"
	"testing"
)

func Test_maxAge(t *testing.T) {
	pp := []interface{}{
		employee{
			Name: "Employee1",
			Age:  18,
		},
		employee{
			Name: "Employee2",
			Age:  20,
		},
		customer{
			Name: "Customer1",
			Age:  34,
		},
		customer{
			Name: "Customer2",
			Age:  65,
		},
	}

	want := pp[3]
	if got := maxAge(pp...); !reflect.DeepEqual(got, want) {
		t.Errorf("MaxAge() = %v, want %v", got, want)
	}
}
