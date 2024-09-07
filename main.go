package main

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"mad_backend_v1/app_middlewares"
	"mad_backend_v1/entities/todos"
	"mad_backend_v1/entities/users"
	"mad_backend_v1/utils/response"
	"net/http"
	"net/url"
	"os"
	"time"
)

type InfoResponse struct {
	Server string `json:"server"`
	Time   int64  `json:"time"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("PORT")
	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresDb := os.Getenv("POSTGRES_DB")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")

	//    ctx := context.Background()

	dsn := url.URL{
		Scheme: "postgres",
		Host:   postgresHost,
		Path:   postgresDb,
		User:   url.UserPassword(postgresUser, postgresPassword),
	}
	//    dsn := "user=postgres password= dbname=altstore_v1 host=localhost port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Printf("failed to connect database: %v", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	//    r.Use(middleware.RealIP)
	r.Use(middleware.AllowContentEncoding("deflate", "gzip"))
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(app_middlewares.JsonMiddleware)
	r.Use(app_middlewares.DBMiddleware(db))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		response.Error[any](w, 404, nil, errors.New("route_does_not_exist"))
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		response.Error[any](w, 405, nil, errors.New("route_method_not_allowed"))
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		infoResponse := InfoResponse{
			Server: "MAD server",
			Time:   time.Now().UnixMilli(),
		}

		response.Success[InfoResponse](w, 200, infoResponse)
	})

	apiRouter := chi.NewRouter()
	v1Router := chi.NewRouter()

	v1Router.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test!!!"))
	})

	v1Router.With(app_middlewares.ProtectedMiddleware).Get("/test2", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test 2!!!"))
	})

	users.GetRouter(v1Router)
	todos.GetRouter(v1Router)

	apiRouter.Mount("/v1", v1Router)
	r.Mount("/api", apiRouter)

	err = http.ListenAndServe(port, r)
	if err != nil {
		fmt.Printf("error happened when listening: %s\n", err)
		return
	}
}
