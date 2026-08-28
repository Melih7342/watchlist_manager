package main

func rcRule(req ScreeningRequest, indexedList IndexedWatchlist) MatchResult {
	hasFirstName := containsColumn(indexedList.ActiveColumns, "firstName")
	hasLastName := containsColumn(indexedList.ActiveColumns, "lastName")
	hasAliases := containsColumn(indexedList.ActiveColumns, "aliases")
	hasDOB := containsColumn(indexedList.ActiveColumns, "dob")

	for _, entry := range indexedList.Entries {

		var firstNameHit, lastNameHit bool
		var firstNameScore, lastNameScore float64

		if hasFirstName {
			firstNameHit, firstNameScore = evaluateFuzzyMatch(entry.FirstName, req.FirstName, 85.0)
		}
		if hasLastName {
			lastNameHit, lastNameScore = evaluateFuzzyMatch(entry.LastName, req.LastName, 85.0)
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
			if highestAliasScore >= 75.0 {
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
				RuleName:      "rcRule",
				WatchlistID:   entry.ID,
				WatchlistName: entry.FirstName + " " + entry.LastName,
				Details:       hitDetails,
			}
		}
	}

	return MatchResult{
		IsHit:    false,
		RuleName: "rcRule",
	}
}
