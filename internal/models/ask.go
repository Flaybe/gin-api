package models

// represents the input for the /ask endpoint
type AskRequest struct {
	Question string `json:"question"`
}

// represents the output for the /ask endpoint
type AskResponse struct {
	Answer string `json:"answer"`
	Source string `json:"source"`
}
