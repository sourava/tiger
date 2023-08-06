package service

import (
	"github.com/sourava/tiger/external/client/sendgrid"
)

type NotificationService struct {
	sendgridClient sendgrid.SendgridApi
}

func NewNotificationService(sendgridClient sendgrid.SendgridApi) *NotificationService {
	return &NotificationService{
		sendgridClient: sendgridClient,
	}
}

func (service *NotificationService) SendMail(name string, email string, subject string, content string) error {
	return service.sendgridClient.SendEmail(&sendgrid.SendEmailRequest{
		RecieverName:  name,
		RecieverEmail: email,
		Subject:       subject,
		Content:       content,
	})
}
