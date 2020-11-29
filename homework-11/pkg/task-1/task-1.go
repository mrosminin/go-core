package task_1

type employee struct {
	Name string
	Age  int
}

type customer struct {
	Name string
	Age  int
}

func (e employee) getAge() int {
	return e.Age
}

func (c customer) getAge() int {
	return c.Age
}

type ageGetter interface {
	getAge() int
}

// maxAge возвращает возраст самого старого покупателя или сотрудника
func maxAge(pp ...ageGetter) int {
	var max int
	for _, p := range pp {
		if age := p.getAge(); age > max {
			max = age
		}
	}
	return max
}
