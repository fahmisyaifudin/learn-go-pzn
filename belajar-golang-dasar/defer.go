package main

import "fmt"

func logging()  {
	fmt.Println("Catat log ...")
}

func runApplication()  {
	defer logging(); //defer akan selalu dieksekusi setelah func selesai, walaupun terjadi error di func yg dieksekusi
	fmt.Println("Application is running")
}

func main()  {
	runApplication()
}