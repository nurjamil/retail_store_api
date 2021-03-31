package middlewares

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func CreateToken(userId int, userRole string) (string, error) {
	claims := jwt.MapClaims{}
	claims["Role"] = userRole
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}

func ExtractTokenUserId(c echo.Context) float64 {
	user := c.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return userId
	}
	return 0
}

func ExtractTokenUserRole(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userRole := claims["Role"].(string)
		return userRole
	}
	return ""
}