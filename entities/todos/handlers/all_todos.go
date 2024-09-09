package handlers

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"mad_backend_v1/app_middlewares"
	"mad_backend_v1/models"
	"mad_backend_v1/utils/database"
	"mad_backend_v1/utils/response"
	"net/http"
)

func allTodos(db *gorm.DB, userId uuid.UUID) ([]models.Todo, error) {
	var todos []models.Todo

	result := db.Table("todos").Joins("join users_todos on users_todos.todo_id = todos.id").Where("users_todos.user_id = ?", userId).Scan(&todos)
	if result.Error != nil {
		pgError := database.ErrorHandler(result.Error)
		if pgError != nil {
			return nil, pgError
		}
	}

	return todos, nil
}

func AllTodosHandler(writer http.ResponseWriter, request *http.Request) {
	db := app_middlewares.GetDBFromContext(request.Context())
	userCtx := app_middlewares.GetUserFromContext(request.Context())

	todos, err := allTodos(db, userCtx.ID)

	if err != nil {
		response.Error[any](writer, 500, nil, err)
		return
	}

	response.Success[[]models.Todo](writer, 200, todos)
	return
}
