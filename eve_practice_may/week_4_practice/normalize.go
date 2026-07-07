package main

import (
	"strings"
)

func NormalizeRows(rows []string) []string {

	if len(rows) == 0 {
		return rows
	}
	var result = make([]string, len(rows))

	for i, ch := range rows {

		result[i] = strings.TrimSpace(strings.ToLower(ch))
	}
	return result
}
