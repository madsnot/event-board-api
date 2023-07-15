package email

import (
	"context"

	"github.com/madsnot/event-board-api/internal/config"
)

type Email struct {
	Address   string
	Password  string
	Host      string
	Port      string
	Recipient string
}

func NewEmailService(cfg config.EmailConfig) *Email {
	return &Email{
		Address:   cfg.Address,
		Password:  cfg.Password,
		Host:      cfg.Host,
		Port:      cfg.Port,
		Recipient: "",
	}
}

func (email *Email) SendMail(code string) error {
	return nil
}

func (email *Email) Verify(ctx context.Context) bool {
	return true
}
