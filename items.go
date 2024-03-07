package main

import "sort"

type Items map[string]int

// Top returns the top n items based on their occurrence count in the Items map.
func (items Items) Top(n int) []string {
	keys := make([]string, 0, len(items))
	for k := range items {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return items[keys[i]] > items[keys[j]]
	})
	return keys[:n]
}
