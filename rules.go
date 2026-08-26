package watchlist_manager

import (
	"encoding/json"
	"os"
)

func loadIndexedList(filePath string, indexes []string) (IndexedWatchlist, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var rawList []WatchlistEntry
	if err := json.Unmarshal(data, &rawList); err != nil {
		return nil, err
	}

	var indexedList []WatchlistEntry

	for _, rawEntry := range rawList {
		indexedEntry := WatchlistEntry{
			ID: rawEntry.ID,
		}

		for _, col := range indexes {
			switch col {
			case "firstName":
				indexedEntry.FirstName = rawEntry.FirstName
			case "lastName":
				indexedEntry.LastName = rawEntry.LastName
			case "aliases":
				indexedEntry.Aliases = rawEntry.Aliases
			case "dob":
				indexedEntry.DOB = rawEntry.DOB
			case "nationality":
				indexedEntry.Nationality = rawEntry.Nationality
			}
		}

		indexedList = append(indexedList, indexedEntry)
	}

	result := IndexedWatchlist{
		Entries: indexedList,
		ActiveColumns: indexes,
	}
	return result, nil
}

func litRule(req ScreeningRequest, indexedList []WatchlistEntry) (bool, float64) {
	for _, indexedEntry := range indexedList {
		if indexedEntry.ID
	}
}
