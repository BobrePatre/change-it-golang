package middlewares

import (
	"change-it/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/sirupsen/logrus"
	"net/http"
)

// TokenAuthMiddleware Middleware to validate JWT tokens
func TokenAuthMiddleware(jwksUrl string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenString := ctx.GetHeader("Authorization")

		// Fetch JWKs from Keycloak
		set, err := jwk.Fetch(ctx, jwksUrl)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch JWKs"})
			ctx.Abort()
			return
		}

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (rawKey interface{}, err error) {
			// Don't forget to validate the alg is what you expect
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				err = fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			keyID, ok := token.Header["kid"].(string)
			if !ok {
				err = fmt.Errorf("expecting JWT header to have string 'kid'")
				return nil, err
			}

			key, found := set.LookupKeyID(keyID)
			if !found {
				err = fmt.Errorf("unable to find key")
				return nil, err
			}

			err = key.Raw(&rawKey)

			return rawKey, nil
		})

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token", "error": err.Error()})
			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Token is valid
			logger.Info("calims", logrus.Fields{"claims": claims})
			ctx.Set("claims", claims)
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

// TODO: доделать
