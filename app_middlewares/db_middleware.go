package app_middlewares

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"mad_backend_v1/utils/response"
	"net/http"
)

type TDBKey string

const dbKey TDBKey = "db"

func DBMiddleware(db *gorm.DB) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			if db == nil {
				response.Error[any](writer, 500, nil, errors.New("database_connection_not_found"))
				return
			}

			ctx := context.WithValue(request.Context(), dbKey, db)
			next.ServeHTTP(writer, request.WithContext(ctx))
		})
	}
}

func GetDBFromContext(ctx context.Context) *gorm.DB {
	db, ok := ctx.Value(dbKey).(*gorm.DB)
	if !ok {
		return nil
	}
	return db
}
