package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		now := time.Now()
		fmt.Printf("%02d:%02d, Hello Women Tech Makers!\n", now.Minute(), now.Second())

		time.Sleep(time.Second)
	}
}
