package response

type CreateUserResponse struct {
	ID       uint   `json:"id" example:"1"`
	Username string `json:"username" example:"username"`
	Email    string `gorm:"unique" example:"user@gmail.com"`
}

type CreateUserHandlerResponse struct {
	Success bool               `json:"success" example:"true"`
	Payload CreateUserResponse `json:"payload"`
}
