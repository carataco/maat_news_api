package authors

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/carataco/maat_news_api/internal/data"
	"github.com/carataco/maat_news_api/internal/request"
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

func GetAuthorsData(r *events.APIGatewayV2HTTPRequest) ([]map[string]any, error) {

	columns := request.ParseColumns(r.QueryStringParameters, defaultAuthorsColumns, allowedAuthorsColumns)

	rows, err := data.SelectAuthors(columns)

	if err != nil {
		return nil, err
	}

	return rows, nil
}
