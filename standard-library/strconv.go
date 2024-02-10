package main

import (
	"fmt"
	"strconv"
)

func main()  {
	resultBool, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error", err.Error())
	}else{
		fmt.Println(resultBool)
	}

	resultInt, _ := strconv.Atoi("10000")
	fmt.Println(resultInt)

	binary := strconv.FormatInt(99, 2)
	fmt.Println(binary)

	stringInt := strconv.Itoa(1000)
	fmt.Println(stringInt)
}