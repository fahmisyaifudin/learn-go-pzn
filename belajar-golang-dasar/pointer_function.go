package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address)  {
	address.Country = "Indonesia"
}

func main()  {
	var address1 Address = Address{
		City: "Lamongan",
		Province: "Jawa Timur",
	}
	fmt.Println(address1)

	ChangeCountryToIndonesia(&address1)

	fmt.Println(address1)
}