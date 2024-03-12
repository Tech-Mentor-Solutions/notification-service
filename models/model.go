package models

// Define a struct to hold the JSON data
type EmailRequest struct {
	To   string `json:"to"`
	Name string `json:"name"`
}

type MeetingRequest struct {
	To        string `json:"to"`
	Name      string `json:"name"`
	Timestamp int64  `json:"timestamp"`
	Url       string `json:"url"`
}

type GroupInvite struct {
	To        string `json:"to"`
	GroupName string `json:"group_name"`
	Sender    string `json:"sender"`
	Receiver  string `json:"receiver"`
	Url       string `json:"url"`
}
