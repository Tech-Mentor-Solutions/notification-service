package main

import (
	"log"
	"net/http"

	emailhandler "github.com/email/handlers"
)

func main() {
	http.HandleFunc("/registration", emailhandler.EmailHandler)
	http.HandleFunc("/meetinglink", emailhandler.MeetingHandler)
	http.HandleFunc("/invitegroup", emailhandler.InvitationHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
