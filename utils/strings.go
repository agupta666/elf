package utils

import "strings"

// SplitAndTrim slices s into all substrings separated by sep and returns a slice of
// the substrings between those separators after trimming spaces
func SplitAndTrim(s, sep string) []string {
	xs := strings.Split(s, sep)
	var result []string

	for _, x := range xs {
		e := strings.TrimSpace(x)
		if len(e) > 0 {
			result = append(result, e)
		}

	}

	return result
}
