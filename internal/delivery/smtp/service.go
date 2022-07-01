package smtp

import (
	"bytes"
	"fmt"
	"github.com/ndt080/schedule-manager-backend/internal/configs"
	"html/template"
	"log"
	"net/smtp"
)

type SmtpService struct {
	auth   smtp.Auth
	config configs.SmtpConfig
}

func NewSmtpService(config configs.SmtpConfig) *SmtpService {
	return &SmtpService{
		config: config,
		auth:   smtp.PlainAuth("", config.Username, config.Password, config.Host),
	}
}

func (service *SmtpService) SendEmail(request SmtpRequest) (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := []byte(fmt.Sprintf("Subject: %s\n%s\n%s", request.Subject, mime, request.Body))
	address := fmt.Sprintf("%s:%s", service.config.Host, service.config.Port)
	from := service.config.Username

	if err := smtp.SendMail(address, service.auth, from, request.To, message); err != nil {
		return false, err
	}

	return true, nil
}

func (service *SmtpService) ParseTemplate(templateFileName string, data interface{}) (string, error) {
	path := fmt.Sprintf("templates/%s", templateFileName)
	log.Println(path)
	t, err := template.ParseFiles(path)
	if err != nil {
		return "", err
	}

	buf := &bytes.Buffer{}
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
