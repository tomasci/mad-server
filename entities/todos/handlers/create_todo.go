package handlers

import (
	"mad_backend_v1/utils/response"
	"net/http"
)

func CreateTodoHandler(writer http.ResponseWriter, request *http.Request) {
	response.Success[string](writer, 200, "created todo")
	return
}
