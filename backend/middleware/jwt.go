package middleware

import (
	"log"
	"os"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var k keyfunc.Keyfunc

func InitializeKeyFunc() {
	var err error
	k, err = keyfunc.NewDefault([]string{os.Getenv("JWKS_URL")})
	if err != nil {
		log.Fatalf("Failed to create a keyfunc.Keyfunc from the server's URL.\nError: %s", err)
	}
}

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("access_token")
		if err != nil {
			return echo.ErrUnauthorized
		}
		token, err := jwt.Parse(cookie.Value, k.Keyfunc)
		if err != nil || !token.Valid {
			return echo.ErrUnauthorized
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return echo.ErrUnauthorized
		}

		c.Set("user", claims)
		return next(c)
	}
}
