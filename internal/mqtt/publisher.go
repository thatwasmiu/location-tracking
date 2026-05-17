package mqtt

import (
	"context"
	"encoding/json"

	mqtt "github.com/eclipse-paho/paho.mqtt.golang"
)

type Publisher struct {
	client mqtt.Client
}

func NewPublisher(client mqtt.Client) *Publisher {
	return &Publisher{client: client}
}

func (p *Publisher) Publish(ctx context.Context, topic string, payload interface{}) error {
	data, err := json.Marshal(payload)

}
