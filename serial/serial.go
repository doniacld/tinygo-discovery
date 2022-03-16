package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		// print the current time
		now := time.Now()
		fmt.Printf("%02d:%02d:%02d, Hello Women Tech Makers!\n", now.Hour(), now.Minute(), now.Second())

		// wait for a second
		time.Sleep(time.Second)
	}
}
