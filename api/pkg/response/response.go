package response

import (
	"encoding/json"
	"net/http"
)

type Response[T any] struct {
	Success    *bool  `json:"success"`
	Message    string `json:"message"`
	Data       T      `json:"data,omitempty"`
	StatusCode *int   `json:"-"`
}

func SendResponse(w http.ResponseWriter, body Response[any], headers map[string]string) {
	if body.Success == nil {
		s := true
		body.Success = &s
	}

	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.Header().Set("Content-Type", "application/json")

	if body.StatusCode == nil {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(*body.StatusCode)
	}

	json.NewEncoder(w).Encode(&body)
}

func SendErrorResponse(w http.ResponseWriter, body Response[any], statusCode int) {
	s := false
	body.Success = &s
	body.StatusCode = &statusCode
	SendResponse(w, body, nil)
}
