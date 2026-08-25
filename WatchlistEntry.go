package watchlist_manager

type WatchlistEntry struct {
	ID          string   `json:"id"`
	EntityType  string   `json:"entityType"`
	FirstName   string   `json:"firstName"`
	LastName    string   `json:"lastName"`
	Aliases     []string `json:"aliases"`
	DOB         string   `json:"dob"`
	Nationality string   `json:"nationality"`
}
