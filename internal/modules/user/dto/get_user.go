package dto

type GetUserByUsernameRequest struct {
	Username string `json:"username" validate:"required"`
}
