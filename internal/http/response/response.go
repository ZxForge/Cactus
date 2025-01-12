package response

import (
	"encoding/json"
	"net/http"
)

type ResponceStatus bool

type responseOkJsonAnswer[T any] struct {
	Success ResponceStatus `json:"success"`
	Type    string         `json:"type"`
	Data    T              `json:"data"`
}

type responseFailAnswer struct {
	Success ResponceStatus `json:"success"`
	Type    string         `json:"type"`
	Message string         `json:"message"`
}

type responseValidationAnswer struct {
	Success ResponceStatus    `json:"success"`
	Type    string            `json:"type"`
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}

func ResponseOKJSON[T any](w http.ResponseWriter, data T) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonOK, _ := json.Marshal(responseOkJsonAnswer[T]{
		Success: true,
		Type:    "data",
		Data:    data,
	})
	w.Write(jsonOK)
}

func ResponseFailJSON(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	jsonFail, _ := json.Marshal(responseFailAnswer{
		Success: false,
		Type:    "fail",
		Message: message,
	})
	w.Write(jsonFail)
}

func ResponseValidationJSON(w http.ResponseWriter, message string, errors map[string]string) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	jsonValidation, _ := json.Marshal(responseValidationAnswer{
		Success: false,
		Type:    "validation",
		Message: message,
		Errors:  errors,
	})
	w.Write(jsonValidation)
}
