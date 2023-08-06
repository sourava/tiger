package service

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/sourava/tiger/business/tiger/request"
	"github.com/sourava/tiger/external/client/sendgrid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_WhenSendgridEmailClientReturnsError_ThenReturnErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockedSendgridApiClient := sendgrid.NewMockSendgridApi(ctrl)
	mockedSendgridApiClient.EXPECT().SendEmail(&sendgrid.SendEmailRequest{
		RecieverEmail: "email@valid.com",
		RecieverName:  "name",
		Subject:       "subject",
		Content:       "content",
	}).Return(errors.New("error while sending mail"))

	service := NewNotificationService(mockedSendgridApiClient)

	err := service.SendMail("name", "email@valid.com", "subject", "content")

	assert.NotNil(t, err)
	assert.Equal(t, "error while sending mail", err.Error())
}

func Test_WhenSendTigerSightingNotificationIsCalledWithMultipleReporters_ThenShouldSendMailForAllReporters(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockedSendgridApiClient := sendgrid.NewMockSendgridApi(ctrl)
	mockedSendgridApiClient.EXPECT().SendEmail(&sendgrid.SendEmailRequest{
		RecieverEmail: "email1@email.com",
		RecieverName:  "name1",
		Subject:       "subject",
		Content:       "message",
	}).Return(nil)
	mockedSendgridApiClient.EXPECT().SendEmail(&sendgrid.SendEmailRequest{
		RecieverEmail: "email2@email.com",
		RecieverName:  "name2",
		Subject:       "subject",
		Content:       "message",
	}).Return(nil)
	mockedSendgridApiClient.EXPECT().SendEmail(&sendgrid.SendEmailRequest{
		RecieverEmail: "email3@email.com",
		RecieverName:  "name3",
		Subject:       "subject",
		Content:       "message",
	}).Return(nil)

	service := NewNotificationService(mockedSendgridApiClient)

	service.SendTigerSightingNotification(&request.SendTigerSightingNotificationRequest{
		Reporters: []*request.TigerSightingReporter{
			{
				Email: "email1@email.com",
				Name:  "name1",
			},
			{
				Email: "email2@email.com",
				Name:  "name2",
			},
			{
				Email: "email3@email.com",
				Name:  "name3",
			},
		},
		Subject: "subject",
		Message: "message",
	})
}
