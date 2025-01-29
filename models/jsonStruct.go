package models

// UserResponse represents the structure of our JSON response
type UserResponse struct {
	Email       string `json:"email"`
	CurrentTime string `json:"current_dateTime"`
	GitHubURL   string `json:"github_url"`
}
