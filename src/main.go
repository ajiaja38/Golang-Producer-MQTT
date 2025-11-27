package main

import (
	"fmt"
	"go-producer-mqtt/src/exception"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	exception.ErrorOnFail(err, "Failed Load .env file")

	HOST := os.Getenv("MQTT_HOST")
	USERNAME := os.Getenv("MQTT_USER")
	PASSWORD := os.Getenv("MQTT_PASSWORD")
	TOPIC := os.Getenv("MQTT_TOPIC")

	opts := mqtt.NewClientOptions()
	opts.AddBroker(HOST)
	opts.SetUsername(USERNAME)
	opts.SetPassword(PASSWORD)
	opts.SetClientID("producer-mqtt")

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to connect to MQTT broker: %v", token.Error())
	}

	defer client.Disconnect(250)

	exception.ErrorOnFail(err, "Failed masrhal struct to json")

	payload := "28134261-6a2b-4d52-9037-32c82fdca08d#00"

	token := client.Publish(TOPIC, 0, false, payload)
	token.Wait()

	if token.Error() != nil {
		exception.ErrorOnFail(token.Error(), fmt.Sprintf("Falied Publish message: %v", token.Error()))
	} else {
		log.Printf("Message Published to topic '%s': %s", TOPIC, payload)
	}

	log.Print("MQTT Producer is running... ✨")
}
