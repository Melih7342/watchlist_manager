package watchlist_manager

func litRule(req ScreeningRequest, indexedList IndexedWatchlist) MatchResult {
	hasFirstName := containsColumn(indexedList.ActiveColumns, "firstName")
	hasLastName := containsColumn(indexedList.ActiveColumns, "lastName")
	hasAliases := containsColumn(indexedList.ActiveColumns, "aliases")

	for _, entry := range indexedList.Entries {

		var firstNameHit, lastNameHit bool
		var firstNameScore, lastNameScore float64

		if hasFirstName {
			firstNameHit, firstNameScore = evaluateFuzzyMatch(entry.FirstName, req.FirstName, 80.0)
		}
		if hasLastName {
			lastNameHit, lastNameScore = evaluateFuzzyMatch(entry.LastName, req.LastName, 80.0)
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

		if (firstNameHit && lastNameHit) || aliasesHit {

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

			return MatchResult{
				IsHit:         true,
				RuleName:      "litRule",
				WatchlistID:   entry.ID,
				WatchlistName: entry.FirstName + " " + entry.LastName,
				Details:       hitDetails,
			}
		}
	}

	return MatchResult{
		IsHit:    false,
		RuleName: "litRule",
	}
}
