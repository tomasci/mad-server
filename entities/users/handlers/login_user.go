package handlers

import (
	"errors"
	"gorm.io/gorm"
	"mad_backend_v1/app_middlewares"
	"mad_backend_v1/models"
	"mad_backend_v1/utils/crypto"
	"mad_backend_v1/utils/database"
	"mad_backend_v1/utils/jwt"
	"mad_backend_v1/utils/response"
	"mad_backend_v1/utils/validation"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=1"`
	Password string `json:"password" validate:"required,min=4"`
}

// Validate is an example of additional validation from RequestBodyMiddleware
// I don't know why would you need this, but here it is,
// and you can do anything you want here
func (data LoginRequest) Validate() validation.ResultValidationErrors {
	result := make(validation.ResultValidationErrors)

	const usernameKey string = "username"
	if data.Username == "89tnx13fhdfh238xejw" {
		result[usernameKey] = append(result[usernameKey], "username_equals_89tnx13fhdfh238xejw")
	}

	// you can even check if password already exists in database ... (haha, no you cannot)

	return result
}

type LoginResponse struct {
	Token string `json:"token"`
}

// LoginUser is a function that accepts username and password,
// selects user from database, validates hash and returns accessToken, refreshToken, status and error
func LoginUser(db *gorm.DB, loginData LoginRequest) (string, string, int, error) {
	username := loginData.Username
	password := loginData.Password

	// search user in database
	var user models.User
	result := db.First(&user, "username = ?", username)

	// handle gorm errors
	if result.Error != nil {
		// handling postgres errors
		pgError := database.ErrorHandler(result.Error)
		if pgError != nil {
			return "", "", 401, pgError
		}
	}

	// compare user hashes
	passwordHash := user.Password
	match, err := crypto.HashValidate(password, passwordHash)
	if err != nil {
		return "", "", 401, errors.New("wrong_username_or_password") // same as before, user doesn't need to know if it is wrong password or not
	}

	if match {
		// create access token
		accessTokenData := make(map[string]interface{})
		accessTokenData["id"] = user.ID
		accessTokenExp, _ := jwt.ExpireInMinutes(1)
		accessTokenString, accessTokenErr := jwt.CreateJWTToken(accessTokenData, accessTokenExp)

		if accessTokenErr != nil {
			return "", "", 500, accessTokenErr
		}

		// create refresh token (currently with the same data, so)
		refreshTokenExp, _ := jwt.ExpireInMonths(3)
		refreshTokenString, refreshTokenErr := jwt.CreateJWTToken(accessTokenData, refreshTokenExp)

		if refreshTokenErr != nil {
			return "", "", 500, refreshTokenErr
		}

		return accessTokenString, refreshTokenString, 200, nil
	} else {
		return "", "", 401, errors.New("wrong_username_or_password") // same thing
	}
}

func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	body := app_middlewares.GetRequestBody[LoginRequest](r.Context())
	db := app_middlewares.GetDBFromContext(r.Context())

	accessToken, refreshToken, status, err := LoginUser(db, body)

	if status == 200 {
		_, refreshTokenExpTime := jwt.ExpireInMonths(3)
		jwt.SetRefreshTokenCookie(w, refreshToken, refreshTokenExpTime)
		response.Success[LoginResponse](w, status, LoginResponse{Token: accessToken})
		return
	} else {
		response.Error[any](w, status, nil, err)
		return
	}
}
