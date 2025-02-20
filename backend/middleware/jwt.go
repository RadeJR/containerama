package middleware

import (
	"log"
	"os"
	"strings"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/RadeJR/containerama/types"
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
		var rawToken string
		cookie, err := c.Cookie("access_token")
		if err != nil {
			authHeader := c.Request().Header["Authorization"]
			if len(authHeader) > 0 {
				rawToken = strings.Split(authHeader[0], "Bearer ")[1]
				if rawToken == "" {
					return echo.ErrUnauthorized
				}
			} else {
				return echo.ErrUnauthorized
			}
		} else {
			rawToken = cookie.Value
		}

		token, err := jwt.ParseWithClaims(rawToken, &types.ZitadelClaims{}, k.Keyfunc)
		if err != nil || !token.Valid {
			return echo.ErrUnauthorized
		}

		claims, ok := token.Claims.(*types.ZitadelClaims)
		if !ok || !token.Valid {
			return echo.ErrUnauthorized
		}

		userID := claims.Subject
		var roles []string

		for k := range claims.Roles {
			roles = append(roles, k)
		}

		c.Set("userID", userID)
		c.Set("roles", roles)
		return next(c)
	}
}
