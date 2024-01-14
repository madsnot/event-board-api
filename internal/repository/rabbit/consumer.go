package rabbit

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
)

type Consumer struct {
	messages <-chan amqp.Delivery
	wait     chan bool
	logger   zerolog.Logger
}

func (c Consumer) SendNotificationToClient(ctx context.Context) {
	for {
		select {
		case msg, ok := <-c.messages:
			if !ok {
				c.wait <- true
			}

			c.logger.Debug().Msg(string(msg.Body))
		case <-c.wait:
		// nothing to do
		case <-ctx.Done():
			return
		}
	}
}
