package mqtt

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Config struct {
	BrokerURL string
	ClientID  string
	Username  string
	Password  string
}

func NewClient(config Config) (mqtt.Client, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(config.BrokerURL)
	opts.SetClientID(config.ClientID)
	opts.SetUsername(config.Username)
	opts.SetPassword(config.Password)

	opts.SetAutoReconnect(true)
	opts.SetConnectRetry(true)

	opts.SetKeepAlive(30 * time.Second)
	opts.SetPingTimeout(10 * time.Second)

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, fmt.Errorf("failed to connect to MQTT broker: %w", token.Error())
	}

	return client, nil
}
