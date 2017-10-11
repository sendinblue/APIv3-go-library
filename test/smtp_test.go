package test

import (
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/sendinblue/APIv3-go-library/client/smtp"
	"github.com/sendinblue/APIv3-go-library/models"
)

func TestSendTransacEmail(t *testing.T) {
	sib := getAPIClient(t)
	params := smtp.NewSendTransacEmailParams()
	email := strfmt.Email(testEmail)
	subject := "Test"
	htmlContent := `Hello <a href="https://example.com/">World</a> !`
	params.SendSMTPEmail = &models.SendSMTPEmail{
		Sender: &models.SendSMTPEmailSender{
			Email: &email,
			Name:  "Test",
		},
		Subject: &subject,
		To: []*models.SendSMTPEmailToItems{
			{
				Email: &email,
				Name:  "Test",
			},
		},
		HTMLContent: &htmlContent,
	}
	res, err := sib.SMTP.SendTransacEmail(params, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(spewConfig.Sdump(res.Payload))
}
