package main

import (
	"container/list"
	"fmt"
)

func main()  {
	data := list.New()
	data.PushBack("Fahmi")
	data.PushBack("Syaifudin")

	head := data.Front()
	headNext := head.Next()

	fmt.Println(head.Value, headNext.Value)

	for e := data.Front(); e != nil; e = e.Next(){
		fmt.Println(e.Value)
	}
}