package service

type EmailCommunicate struct {
	Subject    string `json:"subject"`
	Forgot     string `json:"forgot"`
	CodeActive string `json:"code_active"`
	From       string `json:"from"`
}
