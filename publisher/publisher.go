package publisher

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	// brokerLocalhost is the default localhost broker when launching the client
	brokerLocalhost = "localhost:1883"
	// topic is the topic to publish to
	topic = "topic/secret"
	// msg is the message to publish
	msg = "Hello Gopher!"
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

type Publisher struct {
	Client mqtt.Client
	Token  mqtt.Token
}

func (p Publisher) Connect() {
	// create a new set of options for MQTT client
	options := mqtt.NewClientOptions()
	options.AddBroker(brokerLocalhost)
	options.SetClientID("go_mqtt_example")
	options.SetDefaultPublishHandler(messagePubHandler)
	options.OnConnect = connectHandler
	options.OnConnectionLost = connectionLostHandler

	// create and start a client using the above ClientOptions
	p.Client = mqtt.NewClient(options)
	p.Token = p.Client.Connect()
	if p.Token.Wait() && p.Token.Error() != nil {
		panic(p.Token.Error())
	}
}

func (p Publisher) Publish() {
	// publish a message to the topic
	text := fmt.Sprintf("%s: %s", time.Now(), msg)
	p.Token = p.Client.Publish(topic, 0, false, text)
	p.Token.Wait()
	time.Sleep(time.Second)
}
