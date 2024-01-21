package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main()  {
	alamat1 := &Address{}
	alamat2 := alamat1
	alamat2.City = "Jember"
	fmt.Println(alamat1)
	fmt.Println(alamat2)
	fmt.Println("---------")

	alamat3 := new(Address)
	alamat4 := alamat3
	alamat4.City = "Lamongan"
	fmt.Println(alamat3)
	fmt.Println(alamat4)

}