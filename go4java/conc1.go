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
	sleepAndTalk(0, "Hello")
	sleepAndTalk(1, "DevOxx!")
	sleepAndTalk(2, "What's")
	sleepAndTalk(3, "up?")
}
