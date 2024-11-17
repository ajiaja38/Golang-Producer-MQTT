package main

import (
	"encoding/json"
	"fmt"
	"go-producer-mqtt/src/exception"
	"go-producer-mqtt/src/model"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	exception.ErrorOnFail(err, "Failed Load .env file")

	body := model.UserDao{
		Name:        "M. Aji Perdana",
		Age:         23,
		PhoneNumber: "085695951121",
		Address:     "Lampung",
		Message:     "Message From MQTT",
	}

	HOST := os.Getenv("MQTT_HOST")
	USERNAME := os.Getenv("MQTT_USER")
	PASSWORD := os.Getenv("MQTT_PASSWORD")
	TOPIC := os.Getenv("MQTT_TOPIC")

	opts := mqtt.NewClientOptions()
	opts.AddBroker(HOST)
	opts.SetUsername(USERNAME)
	opts.SetPassword(PASSWORD)
	opts.SetClientID("producer")

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to connect to MQTT broker: %v", token.Error())
	}
	defer client.Disconnect(250)

	jsonBody, err := json.Marshal(body)
	exception.ErrorOnFail(err, "Failed masrhal struct to json")

	token := client.Publish(TOPIC, 0, false, jsonBody)
	token.Wait()

	if token.Error() != nil {
		exception.ErrorOnFail(token.Error(), fmt.Sprintf("Falied Publish message: %v", token.Error()))
	} else {
		log.Printf("Message Published to topic '%s': %s", TOPIC, jsonBody)
	}

	fmt.Println("Mqtt Producer Golang")
}
