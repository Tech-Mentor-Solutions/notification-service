package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/email/models"
)

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read body. Error: ", err.Error())
		http.Error(w, "Failed to read the body", http.StatusBadRequest)
		return
	}

	// Unmarshal the JSON data into the struct
	var emailReq models.EmailRequest
	if err := json.Unmarshal(body, &emailReq); err != nil {
		log.Println("Failed to parse the JSON data. Error: ", err.Error())
		http.Error(w, "Failed to parse JSON data", http.StatusBadRequest)
		return
	}

	// Send the email using SendGrid
	if err := SendRegistration(emailReq.To, emailReq.Name); err != nil {
		http.Error(w, "Failed to send mail", http.StatusInternalServerError)
		log.Println("Failed to send registration email:", err.Error())
		return
	}

	// Respond with a confirmation message
	log.Printf("Email sent to %s", emailReq.To)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Email sent successfully"))
}

// Scheduling meeting
func MeetingHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read body. Error: ", err.Error())
		http.Error(w, "Failed to read the body", http.StatusBadRequest)
		return
	}

	var meetingReq models.MeetingRequest
	if err := json.Unmarshal(body, &meetingReq); err != nil {
		log.Println("Failed to parse JSON data:", err.Error())
		http.Error(w, "Failed to parse JSON data", http.StatusBadRequest)
		return
	}

	// Parse date and time strings into time.Time objects
	date, err := time.Parse("2006-01-02", meetingReq.Date)
	if err != nil {
		log.Println("Failed to parse date:", err.Error())
		http.Error(w, "Failed to parse date", http.StatusBadRequest)
		return
	}

	time, err := time.Parse("15:04:05", meetingReq.Time)
	if err != nil {
		log.Println("Failed to parse time:", err.Error())
		http.Error(w, "Failed to parse time", http.StatusBadRequest)
		return
	}

	if err := SendMeetingLink(meetingReq.To, meetingReq.Name, date, time, meetingReq.Url); err != nil {
		log.Println("Failed to send email. Error: ", err.Error())
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	log.Printf("Email sent to %s", meetingReq.To)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Email sent successfully"))
}

// Group invite
func InvitationHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		log.Println("Failed to read body. Error: ", err.Error())
		return
	}

	var groupInvite models.GroupInvite
	if err := json.Unmarshal(body, &groupInvite); err != nil {
		log.Println("Failed to parse the JSON data. Error: ", err.Error())
		return
	}

	if err := SendGroupInvite(groupInvite.To, groupInvite.GroupName, groupInvite.Receiver, groupInvite.Sender, groupInvite.Url); err != nil {
		log.Println("Failed to send email. Error: ", err.Error())
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	log.Printf("Email sent to %s", groupInvite.To)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Email sent successfully"))

}
