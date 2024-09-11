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

type UpdateTodoRequest struct {
	Title string `json:"title" validate:"required,min=1"`
}

func updateTodo(db *gorm.DB, data UpdateTodoRequest, todoId string, userId uuid.UUID) (*models.Todo, error) {
	tx := db.Begin()

	// select todo
	var todo models.Todo
	//result := tx.Joins("JOIN users_todos on users_todos.todo_id = todos.id").Where("todos.id = ? and users_todos.user_id = ?", todoId, userId).First(&todo)
	// at this point I'm starting to think about switching from gorm to something else, too much magic...
	// of course I have another variant below, but it wasn't like this at first
	result := tx.Table("todos").Joins("JOIN users_todos on users_todos.todo_id = todos.id").Where("todos.id = ? and users_todos.user_id = ? and (deleted_at IS NULL or deleted_at = ?)", todoId, userId, time.Time{}).First(&todo)

	if result.Error != nil {
		pgError := database.ErrorHandler(result.Error)
		if pgError != nil {
			return nil, pgError
		}
	}

	// update todo
	todo.Title = data.Title
	result = tx.Save(&todo) // and here again, save todo, but at what table exactly

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

func UpdateTodoHandler(writer http.ResponseWriter, request *http.Request, todoId string) {
	db := app_middlewares.GetDBFromContext(request.Context())
	body := app_middlewares.GetRequestBody[UpdateTodoRequest](request.Context())
	userCtx := app_middlewares.GetUserFromContext(request.Context())

	todo, err := updateTodo(db, body, todoId, userCtx.ID)

	if err != nil {
		response.Error[any](writer, 500, nil, err)
		return
	}

	response.Success[*models.Todo](writer, 200, todo)
	return
}
