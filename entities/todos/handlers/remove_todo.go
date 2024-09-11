package handlers

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"mad_backend_v1/app_middlewares"
	"mad_backend_v1/models"
	"mad_backend_v1/utils/database"
	"mad_backend_v1/utils/response"
	"net/http"
	"time"
)

// mark and replace content with some "removed_permanently" code
// do not delete row

func removeTodo(db *gorm.DB, todoId string, userId uuid.UUID) (bool, error) {
	tx := db.Begin()

	var todo models.Todo
	result := tx.Table("todos").Joins("JOIN users_todos on users_todos.todo_id = todos.id").Where("todos.id = ? and users_todos.user_id = ?", todoId, userId).First(&todo)

	if result.Error != nil {
		pgError := database.ErrorHandler(result.Error)
		if pgError != nil {
			return false, pgError
		}
	}

	todo.DeletedAt = time.Now()
	result = tx.Save(&todo)

	if result.Error != nil {
		tx.Rollback()

		pgError := database.ErrorHandler(result.Error)
		if pgError != nil {
			return false, pgError
		}
	}

	tx.Commit()
	return true, nil
}

func RemoveTodoHandler(writer http.ResponseWriter, request *http.Request, todoId string) {
	db := app_middlewares.GetDBFromContext(request.Context())
	userCtx := app_middlewares.GetUserFromContext(request.Context())

	_, err := removeTodo(db, todoId, userCtx.ID)
	if err != nil {
		response.Error[any](writer, 500, nil, err)
		return
	}

	response.Success[any](writer, 200, nil)
	return
}
