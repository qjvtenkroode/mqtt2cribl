package mqtt2cribl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func StartReceiving() mqtt.Client {
	var broker = os.Getenv("MQTT2CRIBL_BROKER")
	var clientid = os.Getenv("MQTT2CRIBL_CLIENTID")
	var topics = strings.Split(os.Getenv("MQTT2CRIBL_TOPICS"), ";")
	options := mqtt.NewClientOptions()
	options.AddBroker(broker)
	options.SetClientID(clientid)
	options.SetDefaultPublishHandler(NewMessageHandler())

	client := mqtt.NewClient(options)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for _, topic := range topics {
		token = client.Subscribe(topic, 1, nil)
		token.Wait()
		fmt.Printf("Subscribed to topic %s\n", topic)
	}

	return client
}

func NewMessageHandler() mqtt.MessageHandler {
	return func(client mqtt.Client, message mqtt.Message) {
		msg := map[string]interface{}{}
		if message.Topic() == "dsmr/json" {
			json.Unmarshal([]byte(message.Payload()), &msg)
		} else {
			nameSlice := strings.Split(message.Topic(), "/")
			name := nameSlice[len(nameSlice)-1]
			msg[name] = string(message.Payload())
		}
		msg["topic"] = message.Topic()
		mmsg, _ := json.Marshal(msg)
		msgReader := bytes.NewReader(mmsg)

		body, err := SendToCribl(msgReader)
		if err != nil {
			return
		}

		fmt.Printf("Body: %s\n\n", string(body))
	}
}

func NewConnectHandler(client mqtt.Client) mqtt.OnConnectHandler {
	return func(client mqtt.Client) {
		fmt.Println("Connected")
	}
}

func NewConnectionLostHandler(client mqtt.Client, err error) mqtt.ConnectionLostHandler {
	return func(client mqtt.Client, err error) {
		fmt.Printf("Connection lost: %s\n", err.Error())
	}
}
