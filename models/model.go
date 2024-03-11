package models

// Define a struct to hold the JSON data
type EmailRequest struct {
	To   string `json:"to"`
	Name string `json:"name"`
}

type MeetingRequest struct {
	To   string `json:"to"`
	Name string `json:"name"`
	Date string `json:"date"`
	Time string `json:"time"`
	Url  string `json:"url"`
}

type GroupInvite struct {
	To        string `json:"to"`
	GroupName string `json:"group_name"`
	Sender    string `json:"sender"`
	Receiver  string `json:"receiver"`
	Url       string `json:"url"`
}
