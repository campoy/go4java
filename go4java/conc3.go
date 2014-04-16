package main

import (
	"fmt"
	"time"
)

func sleepAndTalk(secs time.Duration, msg string) {
	time.Sleep(secs * time.Second)
	fmt.Printf("%v ", msg)
}

func main() {
	go sleepAndTalk(0, "Hello")
	go sleepAndTalk(1, "DevOxx!")
	go sleepAndTalk(2, "What's")
	go sleepAndTalk(3, "up?")
	time.Sleep(4 * time.Second)
}
