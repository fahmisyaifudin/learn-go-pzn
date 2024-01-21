package main

import "fmt"

func main()  {
	var names [3]string

	names[0] = "Fahmi"
	names[1] = "Mohamad"
	names[2] = "Syaifudin"

	fmt.Println(names[0])

	//var values = [3]int { 90, 80, 75 }
	var values = [...]int { 90, 80, 75 }

	fmt.Println(values[1])
	fmt.Println(len(values))
}