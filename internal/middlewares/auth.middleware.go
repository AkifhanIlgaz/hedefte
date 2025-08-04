package middlewares

import (
	"net/http"
	"strings"

	"github.com/AkifhanIlgaz/hedefte/pkg/response"
	"github.com/AkifhanIlgaz/hedefte/pkg/token"
	"github.com/gin-gonic/gin"
)

const bearerPrefix = "Bearer "

type AuthMiddleware struct {
	tokenManager *token.Manager
}

func NewAuthMiddleware(tokenManager *token.Manager) *AuthMiddleware {
	return &AuthMiddleware{
		tokenManager: tokenManager,
	}
}

func (m *AuthMiddleware) AccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(http.StatusUnauthorized, "Authorization header is empty", response.WithAbort(c))
			return

		}

		if !strings.HasPrefix(authHeader, bearerPrefix) {
			response.Error(http.StatusUnauthorized, "Authorization header format must be Bearer {token}", response.WithAbort(c))
			return
		}

		token := strings.TrimSpace(authHeader[len(bearerPrefix):])
		if token == "" {
			response.Error(http.StatusUnauthorized, "Token not found", response.WithAbort(c))
			return
		}

		claims, err := m.tokenManager.VerifySupabaseJWT(token)
		if err != nil {
			response.Error(http.StatusUnauthorized, err.Error(), response.WithAbort(c))

			return
		}

		c.Set("uid", claims["sub"])

		c.Next()
	}
}
