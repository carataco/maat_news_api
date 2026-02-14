package handlers

/*
Files: country_handler.go

Role: HTTP handlers for each endpoint.

Responsibilities:

Accepts query parameters (columns, limit, offset, start_date, end_date).

Validates query parameters.

Calls the service layer to execute the request.

Encodes the response as JSON with proper pagination metadata.

Why: Keeps the HTTP-specific logic separate from business/data logic.

This is where you translate HTTP concepts → service calls → JSON response.
*/

import (
	"net/http"

	"github.com/carataco/maat_news_api/internal/httpx"
	"github.com/carataco/maat_news_api/internal/service"
)

func GetAuthors(w http.ResponseWriter, r *http.Request) {

	rows, err := service.GetAuthorsData(r)

	if err != nil {
		httpx.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httpx.WriteJSON(w, http.StatusOK, rows)
}
