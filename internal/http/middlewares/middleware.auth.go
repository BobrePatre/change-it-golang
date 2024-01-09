package middlewares

import (
	V1Domains "change-it/internal/business/domains/v1"
	"change-it/internal/config"
	"change-it/pkg/helpers"
	"change-it/pkg/logger"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func KycloakAuthMiddleware(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "No token provided"})
			ctx.Abort()
			return
		}

		if headerArr := strings.Split(authHeader, " "); len(headerArr) != 2 || headerArr[0] != "Bearer" {
			logger.Error("Invalid token format", nil)
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			ctx.Abort()
			return
		}

		token, err := verifyToken(ctx, strings.Split(authHeader, " ")[1])

		if err != nil {
			logger.Error("cannot verify token", logrus.Fields{"err": err.Error()})
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !(ok && token.Valid) {
			logger.Error("cannot get claims", logrus.Fields{"err": err.Error()})
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			ctx.Abort()
			return
		}

		if claims["sub"] == "" || claims["sub"] == nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			ctx.Abort()
			return
		}

		var userRoles []string
		if resourceAccess, ok := claims["resource_access"].(map[string]interface{}); ok {
			if authClient, ok := resourceAccess[config.AppConfig.AUTHClient].(map[string]interface{}); ok {
				if err := mapstructure.Decode(authClient["roles"], &userRoles); err != nil {
					logger.Error("cannot get user roles", logrus.Fields{"err": err.Error()})
					userRoles = []string{}
				}
			}
		}

		logger.Error("userId", logrus.Fields{"userId": claims["sub"]})

		ctx.Set("userDetails", V1Domains.UserDetails{
			Roles:    userRoles,
			UserId:   claims["sub"].(string),
			Email:    claims["email"].(string),
			Username: claims["name"].(string),
		})

		if len(roles) == 0 {
			ctx.Next()
			return
		}

		if !isUserHaveRoles(roles, userRoles) {
			ctx.Status(http.StatusForbidden)
			ctx.Abort()
		}

		ctx.Next()
	}
}

func isUserHaveRoles(roles []string, userRoles []string) bool {
	for _, role := range roles {
		if helpers.IsArrayContains(userRoles, role) {
			return true
		}
	}
	return false
}

func verifyToken(ctx context.Context, tokenString string) (token *jwt.Token, err error) {

	if err := verifyTokenSession(tokenString); err != nil {
		return nil, err
	}

	set, err := jwk.Fetch(ctx, config.AppConfig.AUTHJwkPublicUri)
	if err != nil {
		return nil, err
	}

	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (rawKey interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			err = fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			return nil, err
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
		if err != nil {
			return nil, fmt.Errorf("invalid token")
		}

		return rawKey, err
	})

	if err != nil {
		return nil, err
	}

	return token, nil

}

func verifyTokenSession(tokenString string) (err error) {
	req, err := http.NewRequest("GET", config.AppConfig.AUTHUserInfoEndpoint, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+tokenString)
	cleint := &http.Client{}
	resp, err := cleint.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("token session is die")
	}

	logger.Info("token session is ok", nil)

	return nil
}
