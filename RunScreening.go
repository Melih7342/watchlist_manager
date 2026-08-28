package main

func RunScreening(req ScreeningRequest, litIndexed IndexedWatchlist, mvIndexed IndexedWatchlist,
	ghIndexed IndexedWatchlist, rcIndexed IndexedWatchlist) (r []MatchResult) {

	var result []MatchResult

	litResult := litRule(req, litIndexed)
	mvResult := mvRule(req, mvIndexed)
	ghResult := ghRule(req, ghIndexed)
	rcResult := rcRule(req, rcIndexed)

	result = append(result, litResult, mvResult, ghResult, rcResult)

	return result
}
