package main

import "fmt"

func main()  {
	person := map[string]string {
		"name": "Fahmi",
		"address": "Lamongan",
	}

	fmt.Println(person["name"]);
	fmt.Println(person)

	book := map[string]string {}
	book["title"] = "GOlang"
	book["author"] = "Fahmi"
	book["ups"] = "Salah"

	delete(book, "ups");

	fmt.Println(book)
}