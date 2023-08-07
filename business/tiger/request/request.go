package request

type CreateTigerRequest struct {
	Name              string  `json:"name" example:"tiger name"`
	DateOfBirth       string  `json:"date_of_birth" example:"2020-01-01"`
	LastSeenTimestamp int     `json:"last_seen_timestamp" example:"1691423085"`
	LastSeenLatitude  float64 `json:"last_seen_latitude" example:"1.1"`
	LastSeenLongitude float64 `json:"last_seen_longitude" example:"10.2"`
}

type ListAllTigerRequest struct {
	Offset   int `json:"offset"`
	PageSize int `json:"pageSize"`
}

type ListAllTigerSightingsRequest struct {
	Offset   int `json:"offset"`
	PageSize int `json:"pageSize"`
	TigerID  int `json:"tigerID"`
}

type TigerSightingReporter struct {
	Email string `json:"email"`
	Name  string `json:"username"`
}

type SendTigerSightingNotificationRequest struct {
	Reporters []*TigerSightingReporter
	Message   string
	Subject   string
}

type CreateTigerSightingRequest struct {
	Timestamp int     `json:"timestamp"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Image     string  `json:"image"`
}
