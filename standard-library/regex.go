package main

import (
	"fmt"
	"regexp"
)

func main()  {
	regex := regexp.MustCompile(`^[a-zA-Z]*$`)
	fmt.Println(regex.MatchString("fahmisyaifudin"))
	fmt.Println(regex.MatchString("fahmi syaifudin"))
	fmt.Println(regex.MatchString("@fahmisyaifudin"))
}