package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Hello, 世界")
		time.Sleep(5 * time.Second)
	}
}
