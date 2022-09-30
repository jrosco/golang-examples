package main

import (
	"fmt"
	"time"
)

func main() {

		TimeTicker := time.NewTicker(3 * time.Second)
		
		tickerChannel := make(chan bool)

		go func() {
			for {
				select {
					case timeticker := <-TimeTicker.C:
						fmt.Println("The time for current is : ", timeticker)
					case <-tickerChannel:
						return
				}
			}
		}()
		
		time.Sleep(6 * time.Second)

		TimeTicker.Stop()

		tickerChannel <- true

		fmt.Println("Time for running ticker is completed")

}
