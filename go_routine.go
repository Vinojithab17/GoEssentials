package main

import (
	"fmt"
	"time"
)

func goMain() {
	// channel := make([]chan bool, 4)
	// channel[0] = make(chan bool)
	// channel[1] = make(chan bool)
	// channel[2] = make(chan bool)
	// channel[3] = make(chan bool)
	done := make(chan bool)
	go greet("Hello", done)
	go greet("Hello Hello", done)
	go greetSlow(done)
	go greet("Hello Hello Hello hello", done)

	// for _, done := range channel {
	// 	<-done
	// }

	for doneChannel := range done {
		fmt.Println(doneChannel)
	}

}

func greetSlow(channel chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello Hello Hello")
	channel <- true
	close(channel)
}

func greet(message string, channel chan bool) {
	fmt.Println(message)
	channel <- true
}
