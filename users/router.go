package users

import (
	"github.com/go-chi/chi/v5"
	"mad_backend_v1/app_middlewares"
	"mad_backend_v1/users/handlers"
)

func GetRouter(api chi.Router) {
	usersRouter := chi.NewRouter()
	usersRouter.With(app_middlewares.RequestBodyMiddleware[handlers.LoginRequest]()).Post("/login", handlers.LoginUserHandler)
	usersRouter.With(app_middlewares.RequestBodyMiddleware[handlers.CreateUserRequest]()).Post("/create", handlers.CreateUserHandler)
	usersRouter.With(app_middlewares.ProtectedMiddleware).Post("/private", handlers.LoginUserHandler)
	api.Mount("/users", usersRouter)
}
