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
	PageSize int `json:"page_size"`
}
