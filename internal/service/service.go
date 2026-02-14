package service

import (
	"net/http"
	"strings"

	"github.com/carataco/maat_news_api/internal/data"
)

func parseColumns(r *http.Request, defaultcols []string, allowedcols map[string]struct{}) []string {
	q := r.URL.Query().Get("columns")

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

func GetData(r *http.Request) ([]map[string]any, error) {

	columns := parseColumns(r, defaultArticlesColumns, allowedArticlesColumns)

	rows, err := data.SelectArticles(columns)

	if err != nil {
		return nil, err
	}

	return rows, nil
}
