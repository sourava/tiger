package response

type LoginResponse struct {
	Token string `json:"token" example:"jwt.token"`
}

type LoginHandlerResponse struct {
	Success bool          `json:"success" example:"true"`
	Payload LoginResponse `json:"payload"`
}
