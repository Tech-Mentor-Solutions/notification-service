package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/email/models"
)

func EmailHandler(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read body. Error: ", err.Error())
		return
	}

	// Unmarshal the JSON data into the struct
	var emailReq models.EmailRequest
	if err := json.Unmarshal(body, &emailReq); err != nil {
		log.Println("Failed to parse the JSON data. Error: ", err.Error())
		return
	}

	// Send the email using SendGrid
	if err := SendSendgrid(emailReq.To, emailReq.Name); err != nil {
		fmt.Fprintf(w, "Email wasn't sent to %s for %s", emailReq.To, emailReq.Name)
		return
	}

	// Respond with a confirmation message
	fmt.Fprintf(w, "Email sent to %s for %s", emailReq.To, emailReq.Name)
	log.Printf("Email sent to %s for %s", emailReq.To, emailReq.Name)
}

// Scheduling meeting
func MeetingHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	var meetingReq models.MeetingRequest
	if err := json.Unmarshal(body, &meetingReq); err != nil {
		http.Error(w, "Failed to parse JSON data", http.StatusBadRequest)
		return
	}

	// Parse date and time strings into time.Time objects
	date, err := time.Parse("2006-01-02", meetingReq.Date)
	if err != nil {
		http.Error(w, "Failed to parse date", http.StatusBadRequest)
		return
	}

	time, err := time.Parse("15:04:05", meetingReq.Time)
	if err != nil {
		http.Error(w, "Failed to parse time", http.StatusBadRequest)
		return
	}

	if err := SendMeetingLink(meetingReq.To, meetingReq.Name, date, time, meetingReq.Url); err != nil {
		fmt.Fprintf(w, "Email wasn't sent to %s for %s", meetingReq.To, meetingReq.Name)
		return
	}

	fmt.Fprintf(w, "Email sent to %s for %s", meetingReq.To, meetingReq.Name)
	log.Printf("Email sent to %s for %s", meetingReq.To, meetingReq.Name)
}

//  Group invite

func InvitationHandler(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	// Unmarshal the JSON data into the struct
	var groupInvite models.GroupInvite
	if err := json.Unmarshal(body, &groupInvite); err != nil {
		http.Error(w, "Failed to parse JSON data", http.StatusBadRequest)
		return
	}

	// Send the email using SendGrid
	SendGroupInvite(groupInvite.To, groupInvite.GroupName, groupInvite.Receiver, groupInvite.Sender, groupInvite.Url)

	// Respond with a confirmation message
	fmt.Fprintf(w, "Email sent to %s for %s", groupInvite.To, groupInvite.Receiver)
}
