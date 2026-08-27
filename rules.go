package watchlist_manager

import (
	"encoding/json"
	"os"
)

func loadIndexedList(filePath string, indexes []string) (IndexedWatchlist, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return IndexedWatchlist{}, err
	}

	var rawList []WatchlistEntry
	if err := json.Unmarshal(data, &rawList); err != nil {
		return IndexedWatchlist{}, err
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
		Entries:       indexedList,
		ActiveColumns: indexes,
	}
	return result, nil
}

func containsColumn(columns []string, colName string) bool {
	for _, c := range columns {
		if c == colName {
			return true
		}
	}
	return false
}

func litRule(req ScreeningRequest, indexedList IndexedWatchlist) (string, bool, float64) {
	hasFirstName := containsColumn(indexedList.ActiveColumns, "firstName")
	hasLastName := containsColumn(indexedList.ActiveColumns, "lastName")
	hasAliases := containsColumn(indexedList.ActiveColumns, "aliases")
	hasDOB := containsColumn(indexedList.ActiveColumns, "DOB") // Achtung: Vorher hattest du "dob" klein geschrieben, Go ist Case-Sensitive!

	// Wir iterieren über die Watchlist
	for _, indexedEntry := range indexedList.Entries {

		// 1. Vornamen / Nachnamen checken
		var firstNameHit, lastNameHit bool
		var firstNameScore, lastNameScore float64

		if hasFirstName {
			firstNameHit, firstNameScore = evaluateFuzzyMatch(indexedEntry.FirstName, req.FirstName, 80.0)
		}
		if hasLastName {
			lastNameHit, lastNameScore = evaluateFuzzyMatch(indexedEntry.LastName, req.LastName, 80.0)
		}

		dobHit := false
		if hasDOB {
			dobHit = evaluateDateMatch(indexedEntry.DOB, req.DOB)
		}

		var aliasesHit bool
		var highestAliasScore float64 = 0.0

		if hasAliases {
			for _, listAlias := range indexedEntry.Aliases {
				for _, reqAlias := range req.Aliases {
					score := fuzzyMatch(listAlias, reqAlias)
					if score > highestAliasScore {
						highestAliasScore = score
					}
				}
			}
			if highestAliasScore >= 80.0 {
				aliasesHit = true
			}
		}

		if (firstNameHit && lastNameHit) || aliasesHit {

			finalScore := max(lastNameScore, highestAliasScore)
			return "litRule", true, finalScore
		}
	}

	return "litRule", false, 0.0
}
