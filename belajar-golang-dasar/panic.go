package main

import "fmt"

func logging()  {
	fmt.Println("Catat log ...")
}

func runApplication(err bool)  {
	defer logging()
	if err {
		panic("Error")
	}
}

func main()  {
	runApplication(true)
}