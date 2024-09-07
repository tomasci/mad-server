package handlers

import (
	"mad_backend_v1/utils/response"
	"net/http"
)

func AllTodosHandler(writer http.ResponseWriter, request *http.Request) {
	response.Success[string](writer, 200, "all todos")
	return
}
