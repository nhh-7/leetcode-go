package hot

import (
	"maps"
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, str := range strs {
		tmp := []byte(str)
		slices.Sort(tmp)
		sortedStr := string(tmp)
		m[sortedStr] = append(m[sortedStr], str)
	}
	return slices.Collect(maps.Values(m))
}
