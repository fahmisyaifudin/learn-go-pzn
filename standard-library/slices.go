package main

import (
	"fmt"
	"slices"
)

func main()  {
	names := []string{"Fahmi", "Syaifudin"}
	values := []int{100, 95, 80, 90}

	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Fahmi"))
	fmt.Println(slices.Index(names, "Fahmi"))
}