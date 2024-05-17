//go:build windows

package utils

import (
	"slices"
	"sort"
	"strings"

	"github.com/prometheus-community/windows_exporter/pkg/types"
)

// ExpandEnabledChildCollectors used by more complex Collectors where user input specifies enabled child Collectors.
// Splits provided child Collectors and deduplicate.
func ExpandEnabledChildCollectors(enabled string) []string {
	result := slices.Compact(strings.Split(enabled, ","))
	// Ensure result is ordered, to prevent test failure
	sort.Strings(result)
	return result
}

func ExpandEnabledCollectors(enabled string) []string {
	expanded := strings.Replace(enabled, types.DefaultCollectorsPlaceholder, types.DefaultCollectors, -1)
	separated := strings.Split(expanded, ",")
	unique := map[string]bool{}
	for _, s := range separated {
		if s != "" {
			unique[s] = true
		}
	}
	result := make([]string, 0, len(unique))
	for s := range unique {
		result = append(result, s)
	}
	return result
}
