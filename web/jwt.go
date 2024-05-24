package web

import (
	"basket-keeper/model"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func JwtMiddleware(next echo.HandlerFunc, mandetory bool) echo.HandlerFunc {
	return func(c echo.Context) error {
		var tokenString = c.Request().Header.Get("Authorization")
		if tokenString == "" {
			if mandetory {
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
			}
			return next(c)
		}
		var token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return model.SigningKey, nil
		})
		if err != nil {
			switch err.(type) {
			case jwt.ValidationError:
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
			default:
				return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
			}
		}
		if !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}
		var claims = token.Claims.(jwt.MapClaims)
		var userID = claims["user_id"].(float64) // type cast to userID
		c.Set("user_id", int64(userID))          // Store user ID in context
		return next(c)
	}
}
