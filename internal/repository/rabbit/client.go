package rabbit

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
)

const (
	CreateRoutingKey = "event.create"
)

type Client struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   string

	Exchange string
}

func NewClient(conn *amqp.Connection, channel *amqp.Channel, queue, exchange string) Client {
	return Client{
		conn:     conn,
		channel:  channel,
		queue:    queue,
		Exchange: exchange,
	}
}

func (c Client) CreateMsg(body []byte) amqp.Publishing {
	return amqp.Publishing{
		ContentType: "text/plain",
		Body:        body,
	}
}

func (c Client) Publish(ctx context.Context, routingKey string, msg amqp.Publishing) error {
	var err error

	c.channel, err = c.conn.Channel()
	if err != nil {
		return err
	}

	return c.channel.PublishWithContext(ctx, c.Exchange, routingKey, false, false, msg)
}

func (c Client) CreateConsumer(queue string, logger zerolog.Logger) (Consumer, error) {
	var err error

	c.channel, err = c.conn.Channel()
	if err != nil {
		return Consumer{}, err
	}

	msgs, err := c.channel.Consume(queue, "", false, true, false, false, nil)
	if err != nil {
		return Consumer{}, err
	}

	return Consumer{
		messages: msgs,
		wait:     make(chan bool),
		logger:   logger,
	}, nil
}
