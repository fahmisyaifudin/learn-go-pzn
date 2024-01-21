package main

import "fmt"

func Empty() any  {
	return 1
}

func main()  {
	// data := Empty().(string)
	// fmt.Println("String", data)

	result := Empty()
	// resultString := result.(string)
	// fmt.Println("String", resultString)
	// resultInt := result.(int)
	// fmt.Println("Integer", resultInt)

	switch value := result.(type) {
		case string:
			fmt.Println("String", value)
		case int:
			fmt.Println("integer", value)	
		default:
			fmt.Println("Unknown")
	}
}