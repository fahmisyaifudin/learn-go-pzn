package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError = errors.New("not found error")
)

func GetById(name string) error  {
	if name == "" {
		return ValidationError
	}
	if name != "fahmi" {
		return NotFoundError
	}

	return nil
}

func main()  {
	err := GetById("")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Validation error")
		}else if errors.Is(err, NotFoundError) {
			fmt.Println("Not found error")
		}else {
			fmt.Println("unknown error")
		}
	}
}