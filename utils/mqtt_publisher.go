// Publish a message to a MQTT broker and wait for the message to be published.
// You can choose to use the online broker or the local one. The message is published to the topic "topic/secret"
// and the payload is the message "1: Hello Gopher!".
// e.g. $ go run mqtt_publisher.go -broker=local
package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	// brokerMosquittoTest is a test server available at https://test.mosquitto.org/
	brokerMosquittoTest = "tcp://test.mosquitto.org:1883"
	// brokerLocalhost is the default localhost broker when launching the client
	brokerLocalhost = "localhost:1883"

	// topic is the topic to publish to
	topic = "topic/secret"
	msg   = "Hello Gopher!"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Message %s received on topic %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectionLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection Lost: %s\n", err.Error())
}

func main() {
	// parse the command line arguments for the broker address
	brokerFlag := flag.String("broker", brokerLocalhost, "The broker you want to publish to can be local or test.")
	flag.Parse()
	broker := parseFlag(*brokerFlag)

	closePublisher()

	// create a new set of options for MQTT client
	options := mqtt.NewClientOptions()
	options.AddBroker(broker)
	options.SetClientID("go_mqtt_example")
	options.SetDefaultPublishHandler(messagePubHandler)
	options.OnConnect = connectHandler
	options.OnConnectionLost = connectionLostHandler

	// create and start a client using the above ClientOptions
	client := mqtt.NewClient(options)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// subscribe to the topic if the broker is the online one
	if broker == brokerMosquittoTest {
		token = client.Subscribe(topic, 1, nil)
		token.Wait()
		fmt.Printf("Subscribed to topic %s\n", topic)
	}

	// publish a message to the topic every second
	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("%d: %s", i, msg)
		token = client.Publish(topic, 0, false, text)
		token.Wait()
		time.Sleep(time.Second)
	}

	// disconnect the client from the broker and wait for the disconnect to finish
	client.Disconnect(100)
}

func parseFlag(broker string) string {
	switch broker {
	case "local": // localhost
		return brokerLocalhost
	case "test": // test.mosquitto.org
		return brokerMosquittoTest
	default:
		return brokerLocalhost
	}
}

// close intercepts a termination signal such as ctrl+c
func closePublisher() {
	c := make(chan os.Signal)
	// listen to SIGTERM signal and notify in the channel
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		// read signal from the channel
		<-c
		fmt.Println("Ciao!")
		// exit the program
		os.Exit(1)
	}()
}
