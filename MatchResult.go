package watchlist_manager

type MatchResult struct {
	IsHit         bool               `json:"isHit"`
	RuleName      string             `json:"ruleName"`
	WatchlistID   string             `json:"watchlistId,omitempty"`
	WatchlistName string             `json:"watchlistName,omitempty"`
	Details       map[string]float64 `json:"details,omitempty"`
}
