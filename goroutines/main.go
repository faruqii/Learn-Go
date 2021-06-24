package main

import (
	"fmt"
	"time"
)

func sendInfo(channel chan<- string, info string) {
	 channel <- info
}

func printInfo(channel <-chan string) {
	fmt.Println(<-channel)
}

func main() {

	informations := make(chan string)

	go sendInfo(informations, "Hello")
	go printInfo(informations)
	time.Sleep(5 * time.Second)
}
