package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

type Animal struct {
	Nama string
}

func (animal Animal) GetName() string  {
	return "Ini " + animal.Nama
}

func (person Person) GetName() string  {
	return person.Name
}

func SayHello(name HasName)  {
	fmt.Println(name.GetName())
}

func main()  {
	person := Person{ Name: "Juleha"}
	animal := Animal{ Nama: "Panda"}
	SayHello(person)
	SayHello(animal)
}