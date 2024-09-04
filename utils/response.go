package utils

import (
	"encoding/json"
	"net/http"
)

type Response[DataType interface{}] struct {
	Status  int      `json:"status"`
	Error   *bool    `json:"error,omitempty"`
	Message *string  `json:"message,omitempty"`
	Data    DataType `json:"data,omitempty"`
}

func MakeResponse[DataType interface{}](w http.ResponseWriter, status int, data DataType) {
	response := Response[DataType]{
		Status:  status,
		Error:   nil,
		Message: nil,
		Data:    data,
	}

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

func MakeErrorResponse[DataType interface{}](w http.ResponseWriter, status int, data DataType, err error) {
	error := true
	errorMsg := err.Error()

	response := Response[any]{
		Status:  status,
		Error:   &error,
		Message: &errorMsg,
		Data:    data,
	}

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
