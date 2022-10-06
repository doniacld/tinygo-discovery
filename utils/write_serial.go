package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/tarm/serial"
)

func main() {
	// pass the port value as flag
	port := flag.String("port", "/dev/cu.usbmodem1401", "The serial port the device is connected to.")
	flag.Parse()

	config := &serial.Config{
		Name:        *port,
		Baud:        9600,
		ReadTimeout: time.Second * 250,
		Size:        8,
	}

	// open the serial port
	stream, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal(err)
	}

	// close properly the stream when ctrl+c is hit
	closePort(stream)

	go func() {
		// output is the string builder for the printed output
		var output strings.Builder
		for {
			// read the buffer from the stream
			buf := make([]byte, 128)
			n, err := stream.Read(buf)
			if err != nil {
				// an EOF may happen and we want to print what's in the output
				if errors.Is(err, io.EOF) {
					log.Printf("From the board: %s", output.String())
					break
				}
				log.Fatalf("failed to read stream %v", err.Error())
			}

			// write the new character
			output.Write(buf[:n])

			// we reached a new line
			if strings.HasSuffix(output.String(), "\n") {
				// let's print the stream from the serial port
				log.Printf("From the board: %s", output.String())
				// reset the builder output for a new string
				output.Reset()
			}
		}
	}()

	// scan the stream from the serial port
	scanner := bufio.NewScanner(os.Stdin)

	// read from the keyboard input
	for scanner.Scan() {
		// write the input from the serial port
		_, err = stream.Write([]byte(scanner.Text()))
		if err != nil {
			log.Fatal(err)
		}

		// writes a new line
		_, err = stream.Write([]byte("\n"))
		if err != nil {
			log.Fatal(err)
		}
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// close intercepts a termination signal such as ctrl+c
func closePort(port *serial.Port) {
	c := make(chan os.Signal)
	// listen to any os interruption, SIGTERM signal and notify in the channel
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		// read signal from the channel
		<-c
		fmt.Println("\nCiao!")
		// close the port properly
		port.Close()
		// exit the program
		os.Exit(1)
	}()
}
