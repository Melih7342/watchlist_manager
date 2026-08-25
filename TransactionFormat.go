package watchlist_manager

type ScreeningRequest struct {
	RequestId   string `json:"request_id"`
	System      string `json:"system"`
	FirstName   string `json:"first-name"`
	LastName    string `json:"last-name"`
	DOB         string `json:"dob"`
	Nationality string `json:"nationality"`
}
