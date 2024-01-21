package main

import "fmt"

type Woman struct {
	Name string
}

func (woman *Woman) isPretty()  {
	woman.Name = woman.Name + " is pretty"
}

func main()  {
	name1 := Woman{
		Name: "Kamila",
	}

	name1.isPretty()

	fmt.Println(name1.Name)
}