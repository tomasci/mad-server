package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"time"
)

func CreateJWTToken(data map[string]interface{}, exp int64) (string, error) {
	jwtSecret := []byte(os.Getenv("JWTSECRET"))
	jwtClaims := jwt.MapClaims{
		"exp": exp,
	}

	for key, value := range data {
		jwtClaims[key] = value
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	jwtTokenString, err := jwtToken.SignedString(jwtSecret)

	if err != nil {
		fmt.Printf("token_signing_error %v\n", err)
		return "", errors.New("token_signing_error")
	}

	return jwtTokenString, nil
}

func SetRefreshTokenCookie(w http.ResponseWriter, refreshToken string, exp time.Time) {
	// creating cookie
	cookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/", // basically any path
		HttpOnly: true,
		Secure:   true, // when https (always?)
		SameSite: http.SameSiteStrictMode,
		Expires:  exp,
	}

	// setting cookie
	http.SetCookie(w, cookie)
}

func ExpireInMinutes(minutes int) (int64, time.Time) {
	period := time.Minute * time.Duration(minutes)
	t := time.Now().Add(period)
	return t.Unix(), t
}

func ExpireInHours(hours int) (int64, time.Time) {
	period := time.Hour * time.Duration(hours)
	t := time.Now().Add(period)
	return t.Unix(), t
}

func ExpireInDays(days int) (int64, time.Time) {
	period := time.Hour * 24 * time.Duration(days)
	t := time.Now().Add(period)
	return t.Unix(), t
}

// ExpireInMonths returns unix timestamp in future after specified number of "months".
// Where each month = 30 days, so 3 months is not 91 or 92 days, it is always 90 days.
func ExpireInMonths(months int) (int64, time.Time) {
	period := time.Hour * 24 * 30 * time.Duration(months)
	t := time.Now().Add(period)
	return t.Unix(), t
}

func ParseToken(token string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}

		jwtSecret := []byte(os.Getenv("JWTSECRET"))

		return jwtSecret, nil
	})

	return parsedToken, err
}

func ValidateToken(rawAccessToken string, rawRefreshToken string) (string, error) {
	_, accessTokenErr := ParseToken(rawAccessToken)
	refreshToken, refreshTokenErr := ParseToken(rawRefreshToken)

	// if refresh token is not valid, user must log in again
	if refreshTokenErr != nil {
		return "", errors.New("token_refresh_token_invalid_token")
	}

	// get refresh token claims
	refreshTokenClaims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("token_refresh_token_invalid_claims")
	}

	// if access token is not valid
	if accessTokenErr != nil {
		// create new token
		accessTokenData := make(map[string]interface{})
		accessTokenData["id"] = refreshTokenClaims["id"]
		accessTokenExp, _ := ExpireInMinutes(1)
		newAccessTokenString, newAccessTokenErr := CreateJWTToken(accessTokenData, accessTokenExp)

		// if failed - error
		if newAccessTokenErr != nil {
			return "", errors.New("token_access_token_create_failed")
		}

		// return new token
		return newAccessTokenString, nil
	}

	// if everything is fine - return old token
	return rawAccessToken, nil
}
