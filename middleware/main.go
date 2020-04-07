package main

import (
	"encoding/json"
	"github.com/raketman/middleware"
	"net/http"
)

var availableClientResolver middleware.AvailableClientResolverContract
var tokenResolver middleware.TokenResolverContract
var clientResolver middleware.ClientResolverContract

type ErrorResponse struct {
    Error  ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	Code            string        `json:"code"`
	Message         string        `json:"message"`
}

func main() {
	availableClientResolver = middleware.DefaultAvailableClientResolver{FilePath:"clients.json"}
	middlewareClient := middleware.Middleware{}


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tokenResolver = middleware.DefaultTokenResolver{Request: r}
		clientResolver = middleware.DefaultClientResolver{AvailableClient: availableClientResolver, Request: r}

		response := middlewareClient.Handle(tokenResolver, clientResolver)

		if response.Status == middleware.StatusSuccess {
			w.Write([]byte(response.Payload))
		} else {
			w.WriteHeader(http.StatusForbidden)

			errorDetail := ErrorDetail{Code:"unauthorized_request", Message: response.Message}
			errorResponse := ErrorResponse{Error: errorDetail}

			result, _ := json.Marshal(errorResponse)

			w.Write(result)
		}

	})

	http.ListenAndServe(":9007", nil)

}