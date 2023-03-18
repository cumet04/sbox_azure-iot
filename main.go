package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	hub := "sbox-iot"
	deviceId := "wsl2"
	hostname := fmt.Sprintf("%s.azure-devices.net", hub)

	sas := generateSasToken(hostname, deviceId, os.Getenv("SIGNIN_KEY"), time.Now().Unix()+60*60)

	opts := mqtt.NewClientOptions()
	opts.SetProtocolVersion(4) // 4 - MQTT 3.1.1
	opts.AddBroker(fmt.Sprintf("ssl://%s:8883", hostname))
	opts.SetClientID(deviceId)
	opts.SetUsername(fmt.Sprintf("%s/%s/api-version=2016-11-14", hostname, deviceId))
	opts.SetPassword(sas)
	c := mqtt.NewClient(opts)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Mqtt error: %s", token.Error())
	}

	token := c.Publish(fmt.Sprintf("devices/%s/messages/events/", deviceId), 0, false, "{\"v\":\"gopher\"}")
	token.Wait()

	c.Disconnect(250)

	fmt.Println("Complete publish")
}

func generateSasToken(hostname string, deviceId string, key string, expires int64) string {
	uri := url.QueryEscape(hostname + "/devices/" + deviceId)

	toSign := uri + "\n" + strconv.FormatInt(expires, 10)

	bkey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		panic(err)
	}
	mac := hmac.New(sha256.New, bkey)
	mac.Write([]byte(toSign))
	sig := mac.Sum(nil)
	hash := url.QueryEscape(base64.StdEncoding.EncodeToString(sig))

	// if (policyName) token += "&skn="+policyName;
	// TODO: policyいる？

	return "SharedAccessSignature sr=" + uri + "&sig=" + hash + "&se=" + strconv.FormatInt(expires, 10)
}
