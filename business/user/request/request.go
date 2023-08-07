package request

type CreateUserRequest struct {
	Username string `json:"username" example:"username"`
	Email    string `json:"email" example:"user@gmail.com"`
	Password string `json:"password" example:"password"`
}
