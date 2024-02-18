package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T)  {
	channel := make(chan string)

	go func ()  {
		time.Sleep(2 * time.Second)
		channel <- "Fahmi Syaifudin"
		fmt.Println("Selesai dikirim")	
	}()

	data := <- channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)

	close(channel)
}

func GiveMeResponse(channel chan string)  {
	time.Sleep(2 * time.Second)
	channel <- "Fahmi Syaifudin"
	fmt.Println("Selesai dikirim")	
}

func TestChannelAsParameter(t *testing.T)  {
	channel := make(chan string)

	go GiveMeResponse(channel)

	data := <- channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)

	close(channel)
}

func OnlyIn(channel chan<- string)  {
	time.Sleep(2 * time.Second)
	channel <- "Fahmi Syaifudin"
}

func OnlyOut(channel <-chan string)  {
	data := <- channel
	fmt.Println(data)
}

func TestChannelInOut(t *testing.T)  {
	channel := make(chan string)
	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
	close(channel)
}