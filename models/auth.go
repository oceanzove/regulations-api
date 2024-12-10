package models

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInOutput struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
