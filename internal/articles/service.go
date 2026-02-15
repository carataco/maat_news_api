package articles

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/carataco/maat_news_api/internal/data"
	"github.com/carataco/maat_news_api/internal/request"
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

func GetArticlesData(r *events.APIGatewayV2HTTPRequest) ([]map[string]any, error) {

	columns := request.ParseColumns(r.QueryStringParameters, defaultArticlesColumns, allowedArticlesColumns)

	rows, err := data.SelectArticles(columns)

	if err != nil {
		return nil, err
	}

	return rows, nil
}
