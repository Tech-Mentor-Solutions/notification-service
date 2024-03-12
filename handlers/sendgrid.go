package handlers

import (
	"log"
	"os"
	"time"

	"github.com/email/models"
	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var (
	sendgridAPIKey         string
	registrationTemplateID string
	meetingTemplateID      string
	invitationTemplateID   string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file: ", err.Error())
	}

	sendgridAPIKey = os.Getenv("SENDGRID_API_KEY")
	registrationTemplateID = os.Getenv("REGISTRATION_TEMPLATE_ID")
	meetingTemplateID = os.Getenv("MEETING_TEMPLATE_ID")
	invitationTemplateID = os.Getenv("INVITATION_TEMPLATE_ID")
}

func SendRegistration(emailReq models.EmailRequest) error {
	from := mail.NewEmail("Mentorship", "kiran@ini8labs.tech")
	recipient := mail.NewEmail("Recipient", emailReq.To)

	// Create a dynamic template message
	message := mail.NewV3Mail()
	message.SetFrom(from)
	message.SetTemplateID(registrationTemplateID)

	// Add recipients
	p := mail.NewPersonalization()
	p.AddTos(recipient)

	// Add personalization to message
	message.AddPersonalizations(p)

	p.SetDynamicTemplateData("name", emailReq.Name)

	client := sendgrid.NewSendClient(sendgridAPIKey)
	_, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// sending meeting link
func SendMeetingLink(meetingReq models.MeetingRequest) error {
	from := mail.NewEmail("Mentorship", "kiran@ini8labs.tech")
	recipient := mail.NewEmail("Recipient", meetingReq.To)

	message := mail.NewV3Mail()
	message.SetFrom(from)
	message.SetTemplateID(meetingTemplateID)

	p := mail.NewPersonalization()
	p.AddTos(recipient)

	message.AddPersonalizations(p)

	t := time.Unix(meetingReq.Timestamp, 0)
	formattedDate := t.Format(time.UnixDate)

	p.SetDynamicTemplateData("name", meetingReq.Name)
	p.SetDynamicTemplateData("date", formattedDate)
	p.SetDynamicTemplateData("url", meetingReq.Url)

	client := sendgrid.NewSendClient(sendgridAPIKey)
	_, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// send group invite
func SendGroupInvite(groupInvite models.GroupInvite) error {
	from := mail.NewEmail("Mentorship", "kiran@ini8labs.tech")
	recipient := mail.NewEmail("Recipient", groupInvite.To)

	message := mail.NewV3Mail()
	message.SetFrom(from)
	message.SetTemplateID(invitationTemplateID)

	p := mail.NewPersonalization()
	p.AddTos(recipient)

	message.AddPersonalizations(p)

	p.SetDynamicTemplateData("receiver", groupInvite.Receiver)
	p.SetDynamicTemplateData("sender", groupInvite.Sender)
	p.SetDynamicTemplateData("groupname", groupInvite.GroupName)
	p.SetDynamicTemplateData("url", groupInvite.Url)

	client := sendgrid.NewSendClient(sendgridAPIKey)
	_, err := client.Send(message)
	if err != nil {
		log.Println("Failed to send message ", err.Error())
		return err
	}
	return nil
}
