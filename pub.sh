#!/bin/bash

sas=$(az iot hub generate-sas-token --cs "$CS" | jq -r .sas)

hub=sbox-iot
device=wsl2

mosquitto_pub -d -q 1 \
  -V mqttv311 \
  -p 8883 \
  -h ${hub}.azure-devices.net \
  -i $device \
  -u "${hub}.azure-devices.net/$device/api-version=2016-11-14" \
  -P "$sas" \
  -t "devices/$device/messages/events/" \
  -m '{"v":"hi"}'
