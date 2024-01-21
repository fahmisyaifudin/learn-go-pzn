package main

import "fmt"


type Customer struct {
	Name, Address string
	Age int
}

func (customer Customer) sayHello(name string)  {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main()  {
	user := Customer{
		Name: "Fahmi",
		Address: "Lamongan",
		Age: 19,
	}

	fahmi := Customer{ "Fahmi", "Lamongan", 20}

	fmt.Println(user.Name)
	fmt.Println(user.Address)

	user.sayHello("Agus")

	fmt.Println(fahmi.Name)
	fmt.Println(fahmi.Age)
}