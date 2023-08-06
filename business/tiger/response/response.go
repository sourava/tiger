package response

type TigerResponse struct {
	ID                uint    `json:"id" example:"1"`
	UserID            uint    `json:"user_id" example:"1"`
	Name              string  `json:"name" example:"tiger name"`
	DateOfBirth       string  `json:"date_of_birth" example:"2020-01-13"`
	LastSeenTimestamp int     `json:"last_seen_timestamp" example:"1691354650"`
	LastSeenLatitude  float64 `json:"last_seen_latitude" example:"1.1"`
	LastSeenLongitude float64 `json:"Last_seen_longitude" example:"10.2"`
}

type ListAllTigersResponse struct {
	Tigers []*TigerResponse `json:"tigers"`
}

type ListAllTigersHandlerResponse struct {
	Success bool                  `json:"success" example:"true"`
	Payload ListAllTigersResponse `json:"payload"`
}

type TigerSightingResponse struct {
	ID        uint    `json:"id" example:"1"`
	TigerID   uint    `json:"tiger_id" example:"1"`
	UserID    uint    `json:"user_id" example:"1"`
	Timestamp int     `json:"timestamp" example:"1691354650"`
	Latitude  float64 `json:"latitude" example:"1.1"`
	Longitude float64 `json:"longitude" example:"10.2"`
}

type ListAllTigerSightingsResponse struct {
	TigerSightings []*TigerSightingResponse `json:"tiger_sightings"`
}

type ListAllTigerSightingsHandlerResponse struct {
	Success bool                          `json:"success" example:"true"`
	Payload ListAllTigerSightingsResponse `json:"payload"`
}
