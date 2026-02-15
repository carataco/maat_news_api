package response

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

// JSON marshals the payload
func Marshal(v any) string {
	b, err := json.Marshal(map[string]any{
		"data": v,
	})
	if err != nil {
		log.Println("error marshaling response:", err)
		return "{}"
	}
	return string(b)
}

// Error marshals an error message
func MarshalError(message string) string {
	b, err := json.Marshal(map[string]any{
		"error": map[string]any{
			"message": message,
		},
	})
	if err != nil {
		log.Println("error marshaling error response:", err)
		return `{"error":{"message":"internal error"}}`
	}
	return string(b)
}

// Success builds an APIGatewayV2HTTPResponse
func Success(v any) events.APIGatewayV2HTTPResponse {
	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: Marshal(v),
	}
}

// Error builds an APIGatewayV2HTTPResponse for errors
func Error(status int, message string) events.APIGatewayV2HTTPResponse {
	return events.APIGatewayV2HTTPResponse{
		StatusCode: status,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: MarshalError(message),
	}
}
