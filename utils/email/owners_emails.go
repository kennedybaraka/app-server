package email

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendVerificationEmail(email string) {
	from := mail.NewEmail("Sender Name", "App")
	to := mail.NewEmail("Recipient Name", email)
	subject := "Test email"
	body := "This is a test email."
	message := mail.NewSingleEmail(from, subject, to, body, "")

	// Set up the SendGrid client and send the email.
	client := sendgrid.NewSendClient("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")
	_, err := client.Send(message)
	if err != nil {
		panic(err)
	}
}
