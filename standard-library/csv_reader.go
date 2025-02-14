package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main()  {
	csvString := "m,fahmi,syaifudin\n" +
		"mohamad,fahmi,syaifudin"

	reader := csv.NewReader(strings.NewReader(csvString))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}
}