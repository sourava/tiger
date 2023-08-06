package response

type TigerResponse struct {
	ID                uint    `json:"id" example:"1"`
	UserID            uint    `json:"user_id" example:"1"`
	Name              string  `json:"name" example:"tiger name"`
	DateOfBirth       string  `json:"dateOfBirth" example:"2020-01-13"`
	LastSeenTimestamp int     `json:"lastSeenTimestamp" example:"1691354650"`
	LastSeenLatitude  float64 `json:"lastSeenLatitude" example:"1.1"`
	LastSeenLongitude float64 `json:"LastSeenLongitude" example:"10.2"`
}

type ListAllTigersResponse struct {
	Tigers []*TigerResponse `json:"tigers"`
}

type ListAllTigersHandlerResponse struct {
	Success bool                  `json:"success" example:"true"`
	Payload ListAllTigersResponse `json:"payload"`
}
