package middleware

import (
	"github.com/NewChakrit/golang_web_development/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// ctx.Abort() // 403
		// read Authorization token from request
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header is missing",
			})
			ctx.Abort()
			return
		}

		// remove bearer from token
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		// validate token with salt
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Config.JwtSaltKey), nil
		})
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			ctx.Abort()
			return
		}

		// Set Header
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok {
			ctx.Set("username", claims["username"])
			ctx.Set("email", claims["email"])
		}

		ctx.Next() // pass
	}
}
