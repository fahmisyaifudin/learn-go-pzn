package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main()  {
	var address1 Address = Address{
		City: "Bandung",
		Province: "Jawa Barat",
		Country: "Indonesia",
	}

	var address2 *Address = &address1
	address2.City = "Cimahi"

	fmt.Println(address1)
	fmt.Println(address2)
}