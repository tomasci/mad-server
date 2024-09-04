package app_middlewares

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"mad_backend_v1/utils"
	"net/http"
	"os"
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

		// second part is token
		tokenString := authHeaderParts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
			}

			jwtSecret := []byte(os.Getenv("JWTSECRET"))

			return jwtSecret, nil
		})

		if err != nil {
			fmt.Printf("invalid_authorization_token %v\n", err)
			//http.Error(w, "invalid_authorization_token", http.StatusUnauthorized)
			utils.MakeErrorResponse[any](w, 401, nil, errors.New("invalid_authorization_token"))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			//http.Error(w, "invalid_token_claims", http.StatusUnauthorized)
			utils.MakeErrorResponse[any](w, 401, nil, errors.New("invalid_token_claims"))
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

		fmt.Printf("protected middleware %v, %v, %v\n", claims["id"], claims["exp"], refreshTokenCookie)

		// todo: common func to read token (must return claims)

		next.ServeHTTP(w, r)
	})
}
