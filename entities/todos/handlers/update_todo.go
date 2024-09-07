package handlers

import (
	"mad_backend_v1/utils/response"
	"net/http"
)

func UpdateTodoHandler(writer http.ResponseWriter, request *http.Request, id string) {
	str := "updated todo, id: " + id
	response.Success[string](writer, 200, str)
	return
}
