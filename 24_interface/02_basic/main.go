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
	d1 := dog{"Sam"}
	c1 := cat{"Mamta"}
	p1 := programmer{"Siba"}
	animals := []animal{d1, &c1, p1}

	var a animal
	fmt.Printf("Type of a is %T, value is %v\n", a, a)

	for _, v := range animals {
		speaker(v)
		fmt.Printf("Type of animal is %T\n", v)
	}
}
