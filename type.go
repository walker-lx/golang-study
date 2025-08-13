package main

type Person interface {
	getName() string
}

type Student struct {
	name string
	age  int
}

// func (stu *Student) getName() string {
// 	return stu.name
// }

type Worker struct {
	name   string
	gender string
}

func (w *Worker) getName() string {
	return w.name
}

func main() {
	var p Person = &Student{
		name: "Tom",
		age:  18,
	}

	var _ Person = (*Student)(nil)
	var _ Person = (*Worker)(nil)

	// fmt.Println(p.getName()) // Tom
}
