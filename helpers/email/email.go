package email

import (
	"errors"
	"github.com/spf13/viper"
	"crypto/tls"
	gomail "gopkg.in/mail.v2"
	"context"
)

func SendEmail(c context.Context, to string, subject string, body string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "daffashafwan.dev@gmail.com")
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)
	dial := gomail.NewDialer("smtp.gmail.com", 587, viper.GetString(`email.email`), viper.GetString(`email.password`))
	dial.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := dial.DialAndSend(msg)
	if err != nil {
		return errors.New("gagal Mengirim Email")
	}

	return errors.New("berhasil Kirim Email")
}
