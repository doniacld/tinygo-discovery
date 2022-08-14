package main

import (
	"fmt"
	"time"

	"tinygo.org/x/drivers/net/mqtt"
)

func main() {
	time.Sleep(3000 * time.Millisecond)
	fmt.Printf("Creating client...")
	opts := mqtt.NewClientOptions()
	fmt.Printf("options created\n")
	opts = opts.AddBroker("localhost:1883")
	opts = opts.SetClientID("tinygoclient")
	fmt.Printf("options created with ClientID\n", opts.ClientID)

	client := mqtt.NewClient(opts)

	fmt.Printf("Connecting to MQTT...\n")
	token := client.Connect()
	fmt.Printf("after connect\n")
	if token.Error() != nil {
		failMessage("Failed to connect to MQTT: " + token.Error().Error())
	}
	if token.Wait() && token.Error() != nil {
		fmt.Printf("Connection failed, token error:", token.Error().Error())
		println("Connection failed, token error:", token.Error().Error())
		return
	}

	for {
		fmt.Printf("Publishing...")
		println("Publishing...")
		token = client.Publish("topic/secret", 0, false, "Hello World")
		token.Wait()
		if token.Error() != nil {
			println("Token error:", token.Error())
			failMessage(token.Error().Error())
		}
		fmt.Printf("Published!")
		println("Published!")
		time.Sleep(2 * time.Second)
		fmt.Printf("next!")

	}

	println("Error: disconnecting MQTT...")
	client.Disconnect(100)

	println("Done.")
}

func failMessage(msg string) {
	for {
		println(msg)
		time.Sleep(1 * time.Second)
	}
}
