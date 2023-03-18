package main

import (
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	sas := "SharedAccessSignature..."
	hub := "sbox-iot"
	device := "wsl2"

	opts := mqtt.NewClientOptions()
	opts.SetProtocolVersion(4) // 4 - MQTT 3.1.1
	opts.AddBroker(fmt.Sprintf("ssl://%s.azure-devices.net:8883", hub))
	opts.SetClientID(device)
	opts.SetUsername(fmt.Sprintf("%s.azure-devices.net/%s/api-version=2016-11-14", hub, device))
	opts.SetPassword(sas)
	c := mqtt.NewClient(opts)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Mqtt error: %s", token.Error())
	}

	token := c.Publish(fmt.Sprintf("devices/%s/messages/events/", device), 0, false, "{\"v\":\"gopher\"}")
	token.Wait()

	c.Disconnect(250)

	fmt.Println("Complete publish")
}
