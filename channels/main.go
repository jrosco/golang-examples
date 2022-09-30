package main

import (
	"fmt"
	"time"
)

var start time.Time
var end time.Time

func SendDataToChannel(ch chan time.Time, t time.Time) {
	ch <- t
}

func main() {

	// create channel
	ch := make(chan time.Time)

	// go routine, send time now to channel
	go SendDataToChannel(ch, time.Now())
	start = <-ch

	// sleep between channels
	time.Sleep(1 * time.Second)

	// go routine, send time now to channel
	go SendDataToChannel(ch, time.Now())
	end = <-ch

	// print time details
	fmt.Println("start time:", start)
	fmt.Println("end time:", end)

	// how much time has elapsed?
	elapsed := end.Sub(start)
	fmt.Println("time elapsed:", elapsed)
}
