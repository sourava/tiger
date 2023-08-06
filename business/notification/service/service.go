package service

import (
	log "github.com/sirupsen/logrus"
	"github.com/sourava/tiger/business/tiger/request"
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

func (service *NotificationService) SendTigerSightingNotification(request *request.SendTigerSightingNotificationRequest) {
	for _, reporter := range request.Reporters {
		err := service.SendMail(reporter.Name, reporter.Email, request.Subject, request.Message)
		if err != nil {
			log.Error(err)
		}
	}
}
