package data

import "time"

func SelectArticles(columns []string) ([]map[string]any, error) {
	all := []map[string]any{
		{
			"id":         1,
			"name":       "Sports",
			"country":    "DE",
			"created_at": time.Now().Add(-24 * time.Hour),
		},
		{
			"id":         2,
			"name":       "Health",
			"country":    "US",
			"created_at": time.Now().Add(-48 * time.Hour),
		},
	}

	var result []map[string]any

	for _, row := range all {
		filtered := make(map[string]any)
		for _, col := range columns {
			filtered[col] = row[col]
		}
		result = append(result, filtered)
	}

	return result, nil
}
