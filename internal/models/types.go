package models

const ContextKeyUser = "user"

type RegisterRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type ProfileResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
