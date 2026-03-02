package hot

import (
	"maps"
	"slices"
)

func groupAnagrams1(strs []string) [][]string {
	m := make(map[string][]string)

	for _, str := range strs {
		tmp := []byte(str)
		slices.Sort(tmp)
		sortedStr := string(tmp)
		m[sortedStr] = append(m[sortedStr], str)
	}
	return slices.Collect(maps.Values(m))
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, str := range strs {
		tmp := []byte(str)
		slices.Sort(tmp)
		m[string(tmp)] = append(m[string(tmp)], str)
	}
	return slices.Collect(maps.Values(m))
}
