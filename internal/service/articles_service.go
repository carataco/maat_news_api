package service

import (
	"net/http"

	"github.com/carataco/maat_news_api/internal/data"
)

var defaultArticlesColumns = []string{
	"name",
	"created_at",
}

var allowedArticlesColumns = map[string]struct{}{
	"id":         {},
	"name":       {},
	"country":    {},
	"created_at": {},
}

func GetArticlesData(r *http.Request) ([]map[string]any, error) {

	columns := parseColumns(r, defaultArticlesColumns, allowedArticlesColumns)

	rows, err := data.SelectArticles(columns)

	if err != nil {
		return nil, err
	}

	return rows, nil
}
