package data

func SelectAuthors(columns []string) ([]map[string]any, error) {
	all := []map[string]any{
		{
			"id":    1,
			"name":  "Alice",
			"email": "alice@example.com",
		},
		{
			"id":    2,
			"name":  "Bob",
			"email": "bob@example.com",
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
