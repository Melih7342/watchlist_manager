package main

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

func fuzzyMatch(s1, s2 string) float64 {
	if len(s1) == 0 && len(s2) == 0 {
		return 100.0
	}

	dist := levenshteinDistance(s1, s2)
	maxLen := max(len(s1), len(s2))

	match := (1 - float64(dist)/float64(maxLen)) * 100

	return match
}

func evaluateFuzzyMatch(s1, s2 string, threshold float64) (bool, float64) {
	score := fuzzyMatch(s1, s2)
	return score >= threshold, score
}

func evaluateDateMatch(dob1, dob2 string) bool {
	if dob1 == "" || dob2 == "" {
		return false
	}
	return dob1 == dob2
}
