package main

import "fmt"

func sayHello(firstName string, lastName string)  {
	fmt.Println("Hello " + firstName + " " + lastName)
}

func getHello(name string) string  {
	return "Hello " + name
}

func getFullname() (string, string)  {
	return "Fahmi", "Syaifudin"
}

func getCompleteName() (firstName string, middleName string, lastName string){
	firstName = "Fahmi"
	middleName = "Mohamad"
	lastName = "Syaifudin"

	return firstName, middleName, lastName
}

func main()  {
	sayHello("Fahmi", "Syaifudin")
	helloName := getHello("Fahmi")
	fmt.Println(helloName)
	
	firstName, lastName := getFullname()
	sayHello(firstName, lastName)
}