package main

import "fmt"

func main()  {
	names := []string{"Kamilatu", "Q", "Hanifah"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	person := map[string]string{
		"name": "Fahmi",
		"company" : "ITS", 
	}

	for _, v := range person {
		fmt.Println(v)
	}
}