//go:generate mockgen --build_flags=--mod=mod -package sendgrid -source sendgrid.go -destination ./mock_sendgrid.go

package sendgrid

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendEmailRequest struct {
	RecieverName  string
	RecieverEmail string
	Subject       string
	Content       string
}

type SendgridApi interface {
	SendEmail(request *SendEmailRequest) error
}

type SendgridApiClient struct {
	apiKey      string
	senderEmail string
	senderName  string
}

func NewSendgridApiClient(apiKey string, senderEmail string, senderName string) SendgridApi {
	return &SendgridApiClient{
		apiKey:      apiKey,
		senderEmail: senderEmail,
		senderName:  senderName,
	}
}

func (c *SendgridApiClient) SendEmail(request *SendEmailRequest) error {
	from := mail.NewEmail(c.senderName, c.senderEmail)
	to := mail.NewEmail(request.RecieverName, request.RecieverEmail)
	message := mail.NewSingleEmail(from, request.Subject, to, request.Content, "")
	client := sendgrid.NewSendClient(c.apiKey)

	_, err := client.Send(message)
	if err != nil {
		return err
	}

	return nil
}
