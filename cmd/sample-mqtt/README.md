# `bright-export-mqtt`

This is a simple MQTT Client which will print instant Electricity usage in Watts

You can build the binary with `go build` and run it.

## Configuration

The only configuration required is provided by environment variables:

* `BRIGHT_USERNAME` - your username for the Bright API.
* `BRIGHT_PASSWORD` - your password for the Bright API.
* `MQTT_TOPIC`      - your topic provided by Bright.

Optionally you can specify the following:

* `MQTT_PORT` - MQTT Broker Port to use
* `MQTT_HOST` - MQTT Broker Hostname

It is worth noting that the client is hardcoded to use TLS, although this is
is easy to replace.

## Sample usage

```
% go build
% export  BRIGHT_USERNAME="..."
% export  BRIGHT_PASSWORD="..."
% export  MQTT_TOPIC="..."
% ./sample-mqtt
DEBU[0000] connecting to broker glowmqtt.energyhive.com:8883 
DEBU[0000] subscribing to SMART/HILD/.../#     
645
644
```
