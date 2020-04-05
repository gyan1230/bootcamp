package main

import (
	"fmt"
	"regexp"
)

func main() {
	r, err := regexp.Compile("[c-x]")
	if err != nil {
		fmt.Println(err)
	}

	f := r.FindString("abc")
	fmt.Println(f)
	xs := r.FindAllString("m1aa11ind", 3)
	fmt.Println(xs)

	f = r.FindString("xyz")
	fmt.Println(f)

	r2, err := regexp.Compile("a.c")
	if err != nil {
		fmt.Println(err)
	}
	xs = r2.FindAllString("a2rererc", 5)
	fmt.Println(xs)
}
