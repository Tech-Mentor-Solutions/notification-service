package main

import (
	"log"
	"net/http"

	"github.com/email/handlers"
)

func main() {
	// Registering handlers for different endpoints
	http.HandleFunc("/registration", handlers.RegistrationHandler)
	http.HandleFunc("/meetinglink", handlers.MeetingHandler)
	http.HandleFunc("/invitegroup", handlers.InvitationHandler)

	// Starting the HTTP server
	log.Println(http.ListenAndServe(":8080", nil))
}
