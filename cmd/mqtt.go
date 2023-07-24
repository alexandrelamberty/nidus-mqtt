package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

var wg = sync.WaitGroup{}

func main() {

	wg.Add(1)

	MQTT.DEBUG = log.New(os.Stdout, "", 0)
	MQTT.ERROR = log.New(os.Stdout, "", 0)

	opts := MQTT.NewClientOptions().AddBroker("mqtt://172.20.0.3:1883").SetClientID("nidus-mqtt")
	opts.AutoReconnect = true
	// opts.SetKeepAlive(60 * time.Second)
	// Set the message callback handler
	opts.SetDefaultPublishHandler(f)
	// opts.SetPingTimeout(1 * time.Second)

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// Subscribe to a topic
	if token := c.Subscribe("testtopic/#", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	time.Sleep(6 * time.Second)

	// Publish a message
	token := c.Publish("testtopic/1", 0, false, "Hello World\n")
	token.Wait()

	wg.Wait()

	// time.Sleep(6 * time.Second)

	// // Unscribe
	// if token := c.Unsubscribe("testtopic/#"); token.Wait() && token.Error() != nil {
	// 	fmt.Println(token.Error())
	// 	os.Exit(1)
	// }

	// Disconnect
	//c.Disconnect(250)
	//time.Sleep(1 * time.Second)
}

func actionCallback(client MQTT.Client, msg MQTT.Message) {
	payload := msg.Payload()
	fmt.Println(payload)
}
