package request

type CreateTigerRequest struct {
	Name              string  `json:"name"`
	DateOfBirth       string  `json:"dateOfBirth"`
	LastSeenTimestamp int     `json:"timestamp"`
	LastSeenLatitude  float64 `json:"latitude"`
	LastSeenLongitude float64 `json:"longitude"`
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
