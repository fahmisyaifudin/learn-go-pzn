package main

import (
	"fmt"
	"time"
)

func main(){
	now := time.Now()
	fmt.Println(now.Local())
	fmt.Println(now.UnixMilli())

	utc := time.Date(2024, time.August, 17, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)

	fmt.Println(utc.Local())
	fmt.Println("Tahun ", utc.Year())
	fmt.Println("Bulan ", utc.Month())
	fmt.Println("Tanggal ", utc.Day())

	formater := "2006-01-02 15:04:05"
	value := "2020-10-10 00:00:00"

	valueTime, err := time.Parse(formater, value)
	if err != nil {
		fmt.Println("Error", err.Error())
	}else{
		fmt.Println(valueTime)
	}

}