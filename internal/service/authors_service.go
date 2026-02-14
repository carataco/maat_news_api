package service

import (
	"net/http"

	"github.com/carataco/maat_news_api/internal/data"
)

var defaultAuthorsColumns = []string{
	"id",
	"name",
}

var allowedAuthorsColumns = map[string]struct{}{
	"id":    {},
	"name":  {},
	"email": {},
}

func GetAuthorsData(r *http.Request) ([]map[string]any, error) {

	columns := parseColumns(r, defaultAuthorsColumns, allowedAuthorsColumns)

	rows, err := data.SelectAuthors(columns)

	if err != nil {
		return nil, err
	}

	return rows, nil
}
