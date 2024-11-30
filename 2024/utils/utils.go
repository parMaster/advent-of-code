package utils

import (
	"os"
	"strings"
)

func ReadLines(file string) []string {
	input, _ := os.ReadFile(file)
	lines := strings.Fields(strings.TrimSpace(string(input)))

	ll := []string{}
	for _, l := range lines {
		ll = append(ll, l)
	}
	return ll
}
