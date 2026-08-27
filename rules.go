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

func litRule(req ScreeningRequest, indexedList IndexedWatchlist) MatchResult {
	hasFirstName := containsColumn(indexedList.ActiveColumns, "firstName")
	hasLastName := containsColumn(indexedList.ActiveColumns, "lastName")
	hasAliases := containsColumn(indexedList.ActiveColumns, "aliases")
	hasDOB := containsColumn(indexedList.ActiveColumns, "DOB")

	for _, entry := range indexedList.Entries {

		var firstNameHit, lastNameHit bool
		var firstNameScore, lastNameScore float64

		if hasFirstName {
			firstNameHit, firstNameScore = evaluateFuzzyMatch(entry.FirstName, req.FirstName, 80.0)
		}
		if hasLastName {
			lastNameHit, lastNameScore = evaluateFuzzyMatch(entry.LastName, req.LastName, 80.0)
		}

		dobHit := false
		if hasDOB {
			dobHit = evaluateDateMatch(entry.DOB, req.DOB)
		}

		var aliasesHit bool
		var highestAliasScore float64 = 0.0

		if hasAliases {
			for _, listAlias := range entry.Aliases {
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

		if ((firstNameHit && lastNameHit) || aliasesHit) || (firstNameHit && lastNameHit && dobHit) {

			hitDetails := make(map[string]float64)

			if firstNameHit {
				hitDetails["firstName"] = firstNameScore
			}
			if lastNameHit {
				hitDetails["lastName"] = lastNameScore
			}
			if aliasesHit {
				hitDetails["aliases"] = highestAliasScore
			}
			if dobHit {
				hitDetails["DOB"] = 100.0
			}

			return MatchResult{
				IsHit:         true,
				RuleName:      "litRule_Standard_Match",
				WatchlistID:   entry.ID,
				WatchlistName: entry.FirstName + " " + entry.LastName,
				Details:       hitDetails,
			}
		}
	}

	return MatchResult{
		IsHit:    false,
		RuleName: "litRule_Standard_Match",
	}
}
