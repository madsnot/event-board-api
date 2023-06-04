package email

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	"net/smtp"

	"github.com/gin-gonic/gin"
)

type Email struct {
	Address   string
	Password  string
	Host      string
	Port      string
	Recipient string
}

func NewEmailService(addr string, pass string, host string, port string) *Email {
	return &Email{
		Address:   addr,
		Password:  pass,
		Host:      host,
		Port:      port,
		Recipient: "",
	}
}

func (email *Email) sendMail(code string) error {
	auth := smtp.PlainAuth("", email.Address, email.Password, email.Host)
	msg := fmt.Sprintf("Subject: Confirm email (Polyhappen.ru)\n%s", code)
	errSendMail := smtp.SendMail(email.Host+email.Port, auth, email.Address, []string{email.Recipient}, []byte(msg))
	return errSendMail
}

func (email *Email) Verify(ctx *gin.Context) bool {
	code := make([]byte, 6)
	_, errCode := rand.Read(code)
	if errCode != nil {
		log.Print("errCode: ", errCode)
		ctx.JSON(http.StatusInternalServerError, gin.H{"response: ": "Internal server error"})
		return false
	}
	// verifyCode := fmt.Sprintf("%x", code)
	// errsendMail := email.sendMail(verifyCode)
	// if errsendMail != nil {
	// 	log.Print("errsendMail: ", errsendMail)
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"response: ": "Internal server error"})
	// 	return false
	// }
	return true
}
