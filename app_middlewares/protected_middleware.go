package app_middlewares

import (
	"errors"
	"fmt"
	"mad_backend_v1/utils"
	mjwt "mad_backend_v1/utils/jwt"
	"net/http"
	"strings"
	"time"
)

func ProtectedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now()
		fmt.Printf("current time %v \n", currentTime)

		// get db connection from context
		db := GetDBFromContext(r.Context())

		// check if it exists
		if db == nil {
			//http.Error(w, "database_connection_not_found", http.StatusInternalServerError)
			utils.MakeErrorResponse[any](w, 500, nil, errors.New("database_connection_not_found"))
			return
		}

		// get token from Authorization header
		authHeader := r.Header.Get("Authorization")
		// split by space
		authHeaderParts := strings.Split(authHeader, " ")

		// check if there are 2 parts after splitting
		fmt.Printf("authHeaderParts len %v, %v\n", authHeader, len(authHeaderParts))

		if len(authHeaderParts) != 2 {
			fmt.Printf("authHeaderParts condiiton", len(authHeaderParts))
			//http.Error(w, "invalid_authorization_header", http.StatusUnauthorized)
			utils.MakeErrorResponse[any](w, 400, nil, errors.New("invalid_authorization_header"))
			return
		}

		// read cookie
		cookies := r.Cookies()
		fmt.Printf("cookies %v\n", cookies)

		refreshTokenCookie, cookieErr := r.Cookie("refresh_token")
		if cookieErr != nil {
			fmt.Printf("client has no cookies %v\n", cookieErr)
			//http.Error(w, "empty_cookies", http.StatusUnauthorized)
			utils.MakeErrorResponse[any](w, 400, nil, errors.New("empty_cookies"))
			return
		}

		// second part is token
		accessTokenString := authHeaderParts[1]
		refreshTokenString := refreshTokenCookie.Value

		newAccessToken, tokenValidationError := mjwt.ValidateToken(accessTokenString, refreshTokenString)

		if tokenValidationError != nil {
			utils.MakeErrorResponse[any](w, 401, nil, tokenValidationError)
			return
		}

		if accessTokenString != newAccessToken {
			w.Header().Set("AccessToken", newAccessToken)
		}

		next.ServeHTTP(w, r)
	})
}
