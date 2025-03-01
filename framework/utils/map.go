package utils

import (
	"sort"
)

func SortMapByValue(m map[string]int) []string {
	// Create a slice to hold the keys
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	// Sort the keys slice based on the values in the map
	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] < m[keys[j]]
	})

	return keys
}
