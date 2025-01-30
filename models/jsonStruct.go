package models

// UserResponse represents the structure of our JSON response
type UserResponse struct {
	Email       string `json:"email"`
	CurrentTime string `json:"current_datetime"`
	GitHubURL   string `json:"github_url"`
}
