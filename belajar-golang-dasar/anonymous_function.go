package main

import "fmt"

type Blacklist = func(string) bool
func registerUser(name string, blacklist Blacklist)  {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	}else {
		fmt.Println("Welcome", name)
	}
}

func main()  {
	registerUser("Fahmi", func(name string) bool {
		if name == "Anjing" {
			return true
		}else{
			return false
		}
	})
}