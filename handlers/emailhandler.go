package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Tech-Mentor-Solutions/notification-service/models"
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

	if emailReq.To == "" {
		log.Println("Sender email id is blank")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Sender email id is blank!!"))
		return
	}
	if emailReq.Name == "" {
		emailReq.Name = "User"
	}

	// Send the email using SendGrid
	if err := SendRegistration(emailReq); err != nil {
		log.Println("Failed to send registration email: ", err.Error())
		http.Error(w, "Failed to send mail", http.StatusInternalServerError)
		return
	}

	// Respond with a confirmation message
	log.Println("Email sent to ", emailReq.To)
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
	if meetingReq.To == "" || meetingReq.Url == "" || meetingReq.Timestamp == 0 {
		if meetingReq.To == "" {
			log.Println("Sender email id is blank")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Sender email id is blank \n"))
			return
		}
		if meetingReq.Url == "" {
			log.Println("URL is blank")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("URL is blank \n"))
			return
		}
		if meetingReq.Timestamp == 0 {
			log.Println("Timestamp is blank")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Timestamp is blank \n"))
			return
		}
	}

	if meetingReq.Name == "" {
		meetingReq.Name = "User"
	}

	if err := SendMeetingLink(meetingReq); err != nil {
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
		log.Println("Failed to read body. Error: ", err.Error())
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}

	var groupInvite models.GroupInvite
	if err := json.Unmarshal(body, &groupInvite); err != nil {
		log.Println("Failed to parse the JSON data. Error: ", err.Error())
		http.Error(w, "Failed to parse the JSON data.", http.StatusBadRequest)
		return
	}
	if groupInvite.To == "" || groupInvite.GroupName == "" || groupInvite.Url == "" {
		if groupInvite.To == "" {
			log.Println("Sender email id is blank")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Sender email id is blank \n"))
			return
		}
		if groupInvite.GroupName == "" {
			log.Println("Group Name is blank")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Group Name is blank \n"))
			return
		}
		if groupInvite.Url == "" {
			log.Println("URL is blank")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("URL is blank \n"))
			return
		}
	}

	if groupInvite.Receiver == "" {
		groupInvite.Receiver = "User"
	}

	if err := SendGroupInvite(groupInvite); err != nil {
		log.Println("Failed to send email. Error: ", err.Error())
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	log.Println("Email sent to ", groupInvite.To)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Email sent successfully"))

}
