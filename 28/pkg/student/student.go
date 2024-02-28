package student

type Students struct {
	Name       string
	Age, Grade int
}

func New(name string, age, grade int) *Students {
	return &Students{
		Name:  name,
		Age:   age,
		Grade: grade,
	}
}
