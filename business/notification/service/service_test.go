package service

import (
	"errors"
	"github.com/golang/mock/gomock"
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
