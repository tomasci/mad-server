package todos

import (
	"github.com/go-chi/chi/v5"
	"mad_backend_v1/app_middlewares"
	"mad_backend_v1/entities/todos/handlers"
	"net/http"
)

func GetRouter(api chi.Router) {
	todosRouter := chi.NewRouter()

	todosRouter.With(app_middlewares.ProtectedMiddleware).Get("/all", handlers.AllTodosHandler)
	todosRouter.With(app_middlewares.ProtectedMiddleware).With(app_middlewares.RequestBodyMiddleware[handlers.CreateTodoRequest]()).Post("/create", handlers.CreateTodoHandler)
	todosRouter.With(app_middlewares.ProtectedMiddleware).Put("/update/{todoId}", func(writer http.ResponseWriter, request *http.Request) {
		todoIdParam := chi.URLParam(request, "todoId")
		handlers.UpdateTodoHandler(writer, request, todoIdParam)
	})
	todosRouter.With(app_middlewares.ProtectedMiddleware).Delete("/remove/{todoId}", func(writer http.ResponseWriter, request *http.Request) {
		todoIdParam := chi.URLParam(request, "todoId")
		handlers.RemoveTodoHandler(writer, request, todoIdParam)
	})

	api.Mount("/todos", todosRouter)
}
