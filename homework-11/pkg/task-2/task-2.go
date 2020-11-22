package task_2

type employee struct {
	Name string
	Age  int
}

type customer struct {
	Name string
	Age  int
}

// MaxAge возвращает объекта самого старого покупателя или сотрудника
func maxAge(pp ...interface{}) interface{} {
	var person interface{}
	var max int
	for _, p := range pp {
		if employee, ok := p.(employee); ok {
			if age := employee.Age; age > max {
				max = age
				person = employee
			}
		}
		if customer, ok := p.(customer); ok {
			if age := customer.Age; age > max {
				max = age
				person = customer
			}
		}
	}
	return person
}
