package main

import (
	"fmt"
	"os"
)

func main()  {
	args := os.Args
	for _, v := range args {
		fmt.Println(v)
	}

	name, err := os.Hostname()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(name)
}