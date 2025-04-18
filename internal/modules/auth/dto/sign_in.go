package dto

type SignInRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignInResponse struct {
	Token string `json:"token"`
}
