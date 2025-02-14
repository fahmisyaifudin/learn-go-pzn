package main

import "fmt"

type Men struct {
	Name string
}

func (men *Men) isPretty()  {
	men.Name = men.Name + " is handsome"
}

func main()  {
	name1 := Men{
		Name: "Fahmi",
	}

	name1.isPretty()

	fmt.Println(name1.Name)
}