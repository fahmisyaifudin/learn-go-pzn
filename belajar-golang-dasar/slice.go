package main

import "fmt"

func main()  {
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice := days[5:]
	fmt.Println(daySlice)

	daySlice[0] = "Sabtu baru"
	daySlice[1] = "Minggu baru"

	fmt.Println(days)

	daySlice2 := append(daySlice, "Libur Baru")
	fmt.Println(daySlice2)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Fahmi"
	newSlice[1] = "Syaifudin"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Mohammad")
	newSlice2[0] = "Pahmi"

	fmt.Println(newSlice)
	fmt.Println(newSlice2)

	fromSLice := days[:]
	toSlice  := make([]string, len(fromSLice), cap(fromSLice))

	copy(toSlice, fromSLice)

	fmt.Println(fromSLice)
	fmt.Println(toSlice)

	iniArray := [...]int{1,2,3,4}
	iniSLice := []int{1,2,3,4}
	fmt.Println(iniArray, iniSLice)
}