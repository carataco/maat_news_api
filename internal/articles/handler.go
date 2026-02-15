package articles

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/carataco/maat_news_api/internal/response"
)

func Handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	rows, err := GetArticlesData(&request)

	if err != nil {

		return response.Error(500, "Problem getting data"), nil
	}

	return response.Success(rows), nil
}
