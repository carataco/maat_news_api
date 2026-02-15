package request

import (
	"strings"
)

func ParseColumns(r map[string]string, defaultcols []string, allowedcols map[string]struct{}) []string {
	q := r["columns"]

	if q == "" {
		return defaultcols
	}

	requested := strings.Split(q, ",")
	var columns []string

	for _, col := range requested {
		col = strings.TrimSpace(col)
		if _, ok := allowedcols[col]; ok {
			columns = append(columns, col)
		}
	}

	if len(columns) == 0 {
		return defaultcols
	}

	return columns
}
