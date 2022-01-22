package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"

	zb "github.com/rk295/bright-golang/mqtt"
)

const (
	mqttHostEnv  = "MQTT_HOST"
	mqttPassEnv  = "MQTT_PASSWORD"
	mqttPortEnv  = "MQTT_PORT"
	mqttTopicEnv = "MQTT_TOPIC"
	mqttUserEnv  = "MQTT_USERNAME"

	mqttDefaultPort = "8883"
)

type mqttConfig struct {
	host  string
	pass  string
	port  string
	topic string
	user  string
}

func main() {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)

	m, err := getMQTTConfig()
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("ssl://%s:%s", m.host, m.port))
	opts.SetUsername(m.user)
	opts.SetPassword(m.pass)
	opts.SetTLSConfig(&tls.Config{})

	logger.Debugf("connecting to broker %s:%s", m.host, m.port)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logger.Error(token.Error())
		os.Exit(1)
	}

	logger.Debugf("subscribing to %s", m.topic)
	var qos byte
	if token := client.Subscribe(m.topic, qos, newMessage); token.Wait() && token.Error() != nil {
		logger.Error(token.Error())
		os.Exit(1)
	}

	<-c
}

func newMessage(c mqtt.Client, m mqtt.Message) {

	var decoded zb.Power

	if err := json.Unmarshal(m.Payload(), &decoded); err != nil {
		fmt.Println("error", err)
	}

	fmt.Println(decoded.ElecMtr.Metering.HistoricalConsumption.InstantaneousDemand)
}

func getMQTTConfig() (mqttConfig, error) {

	var m mqttConfig

	mqttHost := os.Getenv(mqttHostEnv)
	if mqttHost == "" {
		mqttHost = "glowmqtt.energyhive.com"
	}
	m.host = mqttHost

	mqttUser := os.Getenv(mqttUserEnv)
	if mqttUser == "" {
		return m, fmt.Errorf("%s must be set to the username to use to connect to the broker", mqttUserEnv)
	}
	m.user = mqttUser

	mqttPass := os.Getenv(mqttPassEnv)
	if mqttPass == "" {
		return m, fmt.Errorf("the %s variable must be set to the connection password", mqttPassEnv)
	}
	m.pass = mqttPass

	mqttTopic := os.Getenv(mqttTopicEnv)
	if mqttTopic == "" {
		return m, fmt.Errorf("the %s variable must be set to the topic", mqttTopicEnv)
	}
	m.topic = mqttTopic

	mqttPort := os.Getenv(mqttPortEnv)
	if mqttPort == "" {
		mqttPort = mqttDefaultPort
	}
	m.port = mqttPort

	return m, nil
}
