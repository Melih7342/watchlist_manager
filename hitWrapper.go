package watchlist_manager

import "fmt"

func hitMessage(rule string, score float64) string {
	return fmt.Sprintf("Rule %v has caused a hit on these fields: \n"+
		"", rule)

}
