package main

type IndexedWatchlist struct {
	Entries       []WatchlistEntry `json:"entries"`
	ActiveColumns []string         `json:"active_columns"`
}
