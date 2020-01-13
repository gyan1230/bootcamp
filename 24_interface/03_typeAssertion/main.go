package main

import "fmt"

type dog struct {
	Name string
}

type cat struct {
	Name string
}

type programmer struct {
	Name string
}

func (d dog) speak() string {
	return fmt.Sprintf("%v :Woof......", d.Name)
}

func (c *cat) speak() string {
	return fmt.Sprintf("%v :Meawon......", c.Name)
}

func (c *cat) walk() string {
	return fmt.Sprintf("%v :can walk......", c.Name)
}

func (p programmer) speak() string {
	return fmt.Sprintf("%v :Design Pattern...", p.Name)
}

type animal interface {
	speak() string
}

func speaker(a animal) {
	fmt.Println(a.speak())
}

func main() {
	var a animal = &cat{Name: "mamata"}
	speaker(a)
	//error as not implemented by animal interface
	// a.walk()
	// checking if the assertion succeded or not
	c, ok := a.(*cat)
	if ok {
		fmt.Println(c.walk())
	} else {
		fmt.Println("type assertion error")
	}

	a = dog{Name: "sam"}
	switch value := a.(type) {
	case *cat:
		fmt.Printf("%#v is &cat type\n", value)
	case dog:
		fmt.Printf("%#v is dog type\n", value)
		speaker(value)
	case programmer:
		fmt.Printf("%#v is programmer type\n", value)
	}

}
