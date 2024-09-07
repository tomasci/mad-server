package todos

import (
	"github.com/go-chi/chi/v5"
	"mad_backend_v1/app_middlewares"
	"net/http"
)

func GetRouter(api chi.Router) {
	todosRouter := chi.NewRouter()

	todosRouter.With(app_middlewares.ProtectedMiddleware).Get("/all", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("all todos here"))
	})
	todosRouter.With(app_middlewares.ProtectedMiddleware).Post("/create", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("created todo"))
	})
	todosRouter.With(app_middlewares.ProtectedMiddleware).Put("/update/{todoId}", func(writer http.ResponseWriter, request *http.Request) {
		todoIdParam := chi.URLParam(request, "todoId")
		writer.Write([]byte("updated todo by id: " + todoIdParam))
	})
	todosRouter.With(app_middlewares.ProtectedMiddleware).Delete("/remove/{todoId}", func(writer http.ResponseWriter, request *http.Request) {
		todoIdParam := chi.URLParam(request, "todoId")
		writer.Write([]byte("mark todo by id as removed: " + todoIdParam))
		// mark and replace content with some "removed_permanently" code
		// do not delete row
	})

	api.Mount("/todos", todosRouter)
}
