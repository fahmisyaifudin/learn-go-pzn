package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main()  {
	 address1 := Address{
		City: "Bandung",
		Province: "Jawa Barat",
		Country: "Indonesia",
	}

	address2 := &address1
	address2.City = "Cimahi" //berubah untuk city

	fmt.Println(address1)
	fmt.Println(address2)

	fmt.Println("-------------")

	//address2 = &Address{"Surabaya", "Jawa Timur", "Indonesia"}  //copy value
	*address2 = Address{"Lamongan", "Jawa Timur", "Indonesia"} //semua berubah
	fmt.Println(address1)
	fmt.Println(address2)
}