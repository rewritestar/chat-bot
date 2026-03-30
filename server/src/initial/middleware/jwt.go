package middleware

import (
	"net/http"
	"os"
	"strings"

	"chat-bot/src/core/cerror"
	core_values "chat-bot/src/core/values"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Abort()
			err := cerror.CommonError{
				Comment: "token is requried",
				Message: "로그인이 필요합니다.",
			}
			cerror.HandleError(c, http.StatusUnauthorized, err)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.Abort()
			err := cerror.CommonError{
				Comment: "token is not valid",
				Message: "로그인 정보가 유효하지 않습니다. 다시 로그인해주세요.",
			}
			cerror.HandleError(c, http.StatusUnauthorized, err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			workerID, ok := claims[core_values.WorkerIDKey].(float64)
			if ok {
				c.Set(core_values.WorkerIDKey, uint(workerID))
			}
		}
		c.Next()
	}
}
