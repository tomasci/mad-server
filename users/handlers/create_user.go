package handlers

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"mad_backend_v1/app_middlewares"
	"mad_backend_v1/models"
	"mad_backend_v1/utils/crypto"
	"mad_backend_v1/utils/database"
	"mad_backend_v1/utils/response"
	"net/http"
)

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=1"`
	Password string `json:"password" validate:"required,min=4"`
	Email    string `json:"email" validate:"required,min=1"`
}

func createUser(db *gorm.DB, data CreateUserRequest) (bool, error) {
	tx := db.Begin()

	// hashing password
	passwordHash, err := crypto.HashCreate(data.Password)
	if err != nil {
		//log.Fatal(err)
		tx.Commit()
		return false, errors.New("password_hash_create_failed")
	}

	// applying user data to user model
	user := models.User{ID: uuid.New(), Username: data.Username, Password: passwordHash, Email: data.Email}
	// trying to create user
	result := tx.Create(&user)

	// handle gorm errors
	if result.Error != nil {
		tx.Rollback()

		// handling postgres errors
		pgError := database.ErrorHandler(result.Error)
		if pgError != nil {

			return false, pgError
		}
	}

	tx.Commit()

	return true, nil
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	body := app_middlewares.GetRequestBody[CreateUserRequest](r.Context())
	db := app_middlewares.GetDBFromContext(r.Context())

	// trying to create user, returning error if not created
	_, err := createUser(db, body)
	if err != nil {
		response.Error[any](w, 500, nil, err)
		return
	}

	response.Success[any](w, 200, nil)
	return
}
