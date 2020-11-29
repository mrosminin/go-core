package task_1

import "testing"

func Test_maxAge(t *testing.T) {
	pp := []ageGetter{
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
	want := 65
	if got := maxAge(pp...); got != want {
		t.Errorf("MaxAge() = %v, want %v", got, want)
	}

	/* Вопрос на полях. Почему нельзя было вызывать maxAge(ee..., cc...). Пишет cannot use []employee as []ageGetter
	ee := []employee{
		{
			Name: "Employee1",
			Age:  18,
		},
		{
			Name: "Employee2",
			Age:  20,
		},
	}
	cc := []customer{
		{
			Name: "Customer1",
			Age:  34,
		},
		{
			Name: "Customer1",
			Age:  65,
		},
	}*/
}
