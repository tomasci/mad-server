package users

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"mad_backend_v1/app_middlewares"
	handlers2 "mad_backend_v1/entities/users/handlers"
	"net/http"
)

func GetRouter(api chi.Router) {
	usersRouter := chi.NewRouter()

	usersRouter.With(app_middlewares.RequestBodyMiddleware[handlers2.LoginRequest]()).Post("/login", handlers2.LoginUserHandler)
	usersRouter.With(app_middlewares.RequestBodyMiddleware[handlers2.CreateUserRequest]()).Post("/create", handlers2.CreateUserHandler)
	usersRouter.With(app_middlewares.ProtectedMiddleware).Post("/private", func(w http.ResponseWriter, r *http.Request) {
		// just testing
		userCtx := app_middlewares.GetUserFromContext(r.Context())
		fmt.Printf("userCtx %v\n", userCtx)
		if userCtx != nil {
			w.Write([]byte(userCtx.Username))
		}
		w.Write([]byte(""))
	})

	api.Mount("/users", usersRouter)
}
