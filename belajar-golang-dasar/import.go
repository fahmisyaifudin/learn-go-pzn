package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main()  {
	result := helper.SayHelloHelper("Fahmi")
	fmt.Println(result)
	fmt.Println("I like " + helper.Application)
}