package handlers

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendSendgrid(to string, name string) error {
	from := mail.NewEmail("Mentorship", "kiran@ini8labs.tech")
	recipient := mail.NewEmail("Recipient", to)

	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file: ", err.Error())
		return err
	}

	// Access the SendGrid   key from.Error() the environment
	sendgridAPIKey := os.Getenv("SENDGRID_API_KEY")
	templateId := os.Getenv("REGISTRATION_TEMPLATE_ID")

	// Create a dynamic template message
	message := mail.NewV3Mail()
	message.SetFrom(from)

	message.SetTemplateID(templateId)

	// Add recipients
	p := mail.NewPersonalization()
	p.AddTos(recipient)

	// p.Subject = ""

	// Add personalization to message
	message.AddPersonalizations(p)

	p.SetDynamicTemplateData("name", name)

	client := sendgrid.NewSendClient(sendgridAPIKey)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(response.Body)
	log.Println(response.StatusCode)
	log.Println(response.Headers)
	return nil
}

// sending meeting link
func SendMeetingLink(to string, name string, date time.Time, time time.Time, url string) error {
	from := mail.NewEmail("Mentorship", "kiran@ini8labs.tech")
	recipient := mail.NewEmail("Recipient", to)

	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file: ", err.Error())
		return err
	}

	sendgridAPIKey := os.Getenv("SENDGRID_API_KEY")
	templateId := os.Getenv("MEETING_TEMPLATE_ID")

	message := mail.NewV3Mail()
	message.SetFrom(from)

	message.SetTemplateID(templateId)

	p := mail.NewPersonalization()
	p.AddTos(recipient)

	// Set subject (you can add your own subject if needed)
	// p.Subject = ""

	// Add personalization to message
	message.AddPersonalizations(p)

	p.SetDynamicTemplateData("name", name)
	p.SetDynamicTemplateData("date", date.Format("2006-01-02"))
	p.SetDynamicTemplateData("time", time.Format("15:04:05"))
	p.SetDynamicTemplateData("url", url)

	client := sendgrid.NewSendClient(sendgridAPIKey)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println(response.StatusCode)
	fmt.Println(response.Body)
	fmt.Println(response.Headers)
	return nil
}

// send group invite
func SendGroupInvite(to string, groupName string, receiver string, sender string, url string) error {
	from := mail.NewEmail("Mentorship", "kiran@ini8labs.tech")
	recipient := mail.NewEmail("Recipient", to)

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	sendgridAPIKey := os.Getenv("SENDGRID_API_KEY")
	templateId := os.Getenv("INVITATION_TEMPLATE_ID")

	message := mail.NewV3Mail()
	message.SetFrom(from)

	message.SetTemplateID(templateId)

	p := mail.NewPersonalization()
	p.AddTos(recipient)

	// Set subject (you can add your own subject if needed)
	// p.Subject = ""

	message.AddPersonalizations(p)

	p.SetDynamicTemplateData("receiver", receiver)
	p.SetDynamicTemplateData("sender", sender)
	p.SetDynamicTemplateData("groupname", groupName)
	p.SetDynamicTemplateData("url", url)

	client := sendgrid.NewSendClient(sendgridAPIKey)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println(response.StatusCode)
	fmt.Println(response.Body)
	fmt.Println(response.Headers)
	return nil
}
