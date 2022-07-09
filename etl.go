package etl

import (
	"strings"
)

func Transform(in map[int][]string) map[string]int {
	result := map[string]int{}
	for i, v := range in {
		for _, vv := range v {
			result[strings.ToLower(vv)] += i
		}
	}
	return result
}
