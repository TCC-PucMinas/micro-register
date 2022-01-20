package service

type EmailCommunicate struct {
	Subject string `json:"subject"`
	Forgot  string `json:"forgot"`
	From    string `json:"from"`
}
