package middlewares

import (
	"change-it/internal/config"
	"change-it/pkg/helpers"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/lestrrat-go/jwx/jwk"
	"net/http"
	"strings"
)

func AuthMiddleware(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "No token provided"})
			ctx.Abort()
			return
		}

		if len(strings.Split(authHeader, " ")) != 2 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			ctx.Abort()
			return
		}

		if strings.Split(authHeader, " ")[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			ctx.Abort()
			return
		}

		token, err := verifyToken(ctx, strings.Split(authHeader, " ")[1])

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			ctx.Set("claims", claims)
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			ctx.Abort()
			return
		}

		userRoles, ok := claims["resource_access"].(map[string]interface{})[config.AppConfig.AUTHClient].(map[string]interface{})["roles"].([]string)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			ctx.Abort()
			return
		}

		for _, role := range roles {
			if helpers.IsArrayContains(userRoles, role) {
				ctx.Next()
				return
			}
		}

		ctx.Status(http.StatusForbidden)
		ctx.Abort()
	}
}

func verifyToken(ctx context.Context, tokenString string) (token *jwt.Token, err error) {
	set, err := jwk.Fetch(ctx, config.AppConfig.AUTHJwkPublicUri)
	if err != nil {
		return nil, err
	}

	// Parse the token
	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (rawKey interface{}, err error) {

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
		return nil, err
	}

	return token, nil

}
