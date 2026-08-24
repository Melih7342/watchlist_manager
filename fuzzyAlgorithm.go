package watchlist_manager

func levenshteinDistance(s1, s2 string) int {
	matrix := make([][]int, len(s1)+1)
	for i := range matrix {
		matrix[i] = make([]int, len(s2)+1)
	}

	for i := 0; i <= len(s1); i++ {
		matrix[i][0] = i
	}
	for j := 0; j <= len(s1); j++ {
		matrix[0][j] = j
	}

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				matrix[i][j] = matrix[i-1][j-1]
			} else {
				insertCost := matrix[i][j-1]
				deleteCost := matrix[i-1][j]
				replaceCost := matrix[i-1][j-1]

				matrix[i][j] = min(insertCost, deleteCost, replaceCost) + 1
			}
		}
	}

	return matrix[len(s1)][len(s2)]
}
