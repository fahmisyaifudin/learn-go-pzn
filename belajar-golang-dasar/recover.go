package main

import "fmt"

func endApp()  {
	fmt.Println("End application...")
	message := recover() //recover ditaruh di defer supaya ketika terjadi error/panic aplikasi tidak benar" terhenti
	fmt.Println("Terjadi error ", message)
}

func runApplication(err bool)  {
	defer endApp()
	if err {
		panic("Ups Error" )
	}
}

func main()  {
	runApplication(true)
	fmt.Println("Hello Fahmi")
}