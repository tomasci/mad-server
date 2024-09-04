package handlers

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"mad_backend_v1/app_middlewares"
	"mad_backend_v1/models"
	"mad_backend_v1/utils"
	"mad_backend_v1/utils/crypto"
	"mad_backend_v1/utils/database"
	"net/http"
)

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=1"`
	Password string `json:"password" validate:"required,min=4"`
	Email    string `json:"email" validate:"required,min=1"`
}

type CreateUserResponse struct {
	Token string `json:"token"`
}

func createUser(db *gorm.DB, data CreateUserRequest) (bool, error) {
	// hashing password
	passwordHash, err := crypto.HashCreate(data.Password)
	if err != nil {
		log.Fatal(err)
	}

	// applying user data to user model
	user := models.User{ID: uuid.New(), Username: data.Username, Password: passwordHash, Email: data.Email}
	// trying to create user
	result := db.Create(&user)

	// handle gorm errors
	if result.Error != nil {
		// handling postgres errors
		pgError := database.ErrorHandler(result.Error)
		if pgError != nil {
			return false, pgError
		}
	}

	return true, nil
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	body := app_middlewares.GetRequestBody[CreateUserRequest](r.Context())
	db := app_middlewares.GetDBFromContext(r.Context())

	// trying to create user, returning error if not created
	_, err := createUser(db, body)
	if err != nil {
		utils.MakeErrorResponse[any](w, 500, nil, err)
		return
	}

	// signing in if user created
	// basically body is the same and contains required username & password
	LoginUserHandler(w, r)
}
