package main

import (
	"fmt"
	"time"
)

func main()  {
	duration1 := time.Second * 100
	duration2 := time.Hour * 1
	duration3 := duration2 - duration1

	fmt.Printf("duration : %d\n", duration3)
}