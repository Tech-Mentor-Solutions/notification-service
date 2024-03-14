package handlers

import (
	"log"
	"os"
	"time"

	"github.com/Tech-Mentor-Solutions/notification-service/models"
	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file: ", err.Error())
	}
}

func SendRegistration(emailReq models.EmailRequest) error {
	from := mail.NewEmail("Mentorship", os.Getenv("FROM"))
	recipient := mail.NewEmail("Recipient", emailReq.To)

	// Create a dynamic template message
	message := mail.NewV3Mail()
	message.SetFrom(from)
	message.SetTemplateID(os.Getenv("REGISTRATION_TEMPLATE_ID"))

	// Add recipients
	p := mail.NewPersonalization()
	p.AddTos(recipient)

	// Add personalization to message
	message.AddPersonalizations(p)

	p.SetDynamicTemplateData("name", emailReq.Name)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	_, err := client.Send(message)
	if err != nil {
		log.Println("Failed to send email ", err.Error())
		return err
	}
	return nil
}

// sending meeting link
func SendMeetingLink(meetingReq models.MeetingRequest) error {
	from := mail.NewEmail("Mentorship", os.Getenv("FROM"))
	recipient := mail.NewEmail("Recipient", meetingReq.To)

	message := mail.NewV3Mail()
	message.SetFrom(from)
	message.SetTemplateID(os.Getenv("MEETING_TEMPLATE_ID"))

	p := mail.NewPersonalization()
	p.AddTos(recipient)

	message.AddPersonalizations(p)

	t := time.Unix(meetingReq.Timestamp, 0)
	formattedDate := t.Format(time.UnixDate)

	p.SetDynamicTemplateData("name", meetingReq.Name)
	p.SetDynamicTemplateData("date", formattedDate)
	p.SetDynamicTemplateData("url", meetingReq.Url)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	_, err := client.Send(message)
	if err != nil {
		log.Println("Failed to send email ", err.Error())
		return err
	}

	return nil
}

// send group invite
func SendGroupInvite(groupInvite models.GroupInvite) error {
	from := mail.NewEmail("Mentorship", os.Getenv("FROM"))
	recipient := mail.NewEmail("Recipient", groupInvite.To)

	message := mail.NewV3Mail()
	message.SetFrom(from)
	message.SetTemplateID(os.Getenv("INVITATION_TEMPLATE_ID"))

	p := mail.NewPersonalization()
	p.AddTos(recipient)

	message.AddPersonalizations(p)

	p.SetDynamicTemplateData("receiver", groupInvite.Receiver)
	p.SetDynamicTemplateData("name", groupInvite.Receiver)
	p.SetDynamicTemplateData("groupname", groupInvite.GroupName)
	p.SetDynamicTemplateData("url", groupInvite.Url)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	_, err := client.Send(message)
	if err != nil {
		log.Println("Failed to send message ", err.Error())
		return err
	}
	return nil
}
