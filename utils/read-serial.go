package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tarm/serial"
)

func main() {
	config := &serial.Config{
		Name:        "/dev/cu.usbmodem1101",
		Baud:        9600,
		ReadTimeout: time.Second * 250,
		Size:        8,
	}

	stream, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal(err)
	}

	// close properly the stream when ctrl+c is hit
	close(stream)

	scanner := bufio.NewScanner(stream)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second * 5)

}

func close(port *serial.Port) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Ciao!")
		port.Close()
		os.Exit(1)
	}()
}
