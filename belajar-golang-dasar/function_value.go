package main

import "fmt"

func getGoodBye(name string) string  {
	return "Good Bye " + name
}

func main()  {
	myFunction	:= getGoodBye

	fmt.Println(myFunction("Fahmi"))
}