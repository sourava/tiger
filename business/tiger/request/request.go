package request

type CreateTigerRequest struct {
	Name              string  `json:"name"`
	DateOfBirth       string  `json:"dateOfBirth"`
	LastSeenTimestamp int     `json:"timestamp"`
	LastSeenLatitude  float64 `json:"latitude"`
	LastSeenLongitude float64 `json:"longitude"`
}
