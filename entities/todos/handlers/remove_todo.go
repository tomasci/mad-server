package handlers

import (
	"mad_backend_v1/utils/response"
	"net/http"
)

// mark and replace content with some "removed_permanently" code
// do not delete row

func RemoveTodoHandler(writer http.ResponseWriter, request *http.Request, id string) {
	str := "removed todo, id: " + id
	response.Success[string](writer, 200, str)
	return
}
