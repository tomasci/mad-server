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

func createTodo(db *gorm.DB, userId uuid.UUID) (*models.Todo, error) {
	tx := db.Begin()

	// create todo
	todo := models.Todo{ID: uuid.New()}
	result := tx.Create(&todo)
	if result.Error != nil {
		tx.Rollback()

		pgError := database.ErrorHandler(result.Error)
		if pgError != nil {
			return nil, pgError
		}
	}

	// create relation
	rel := models.UsersTodos{UserID: userId, TodoID: todo.ID}
	result = tx.Create(&rel)
	if result.Error != nil {
		tx.Rollback()

		pgError := database.ErrorHandler(result.Error)
		if pgError != nil {
			return nil, pgError
		}
	}

	tx.Commit()

	return &todo, nil
}

func CreateTodoHandler(writer http.ResponseWriter, request *http.Request) {
	db := app_middlewares.GetDBFromContext(request.Context())
	userCtx := app_middlewares.GetUserFromContext(request.Context())

	todo, err := createTodo(db, userCtx.ID)
	if err != nil {
		response.Error[any](writer, 500, nil, err)
		return
	}

	response.Success[*models.Todo](writer, 200, todo)
	return
}
