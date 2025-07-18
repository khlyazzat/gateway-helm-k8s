package middleware

import (
	"main/services/api-gateway/internal/values"
	"strings"

	"github.com/gin-gonic/gin"

	"main/pkg/jwt"
)

type Middleware interface {
	Authorize(c *gin.Context)
	Panic(c *gin.Context)
	Error(c *gin.Context)
}

type middleware struct {
	jwt jwt.JWT
}

// Error implements Middleware.
func (m *middleware) Error(c *gin.Context) {
	panic("unimplemented")
}

// Panic implements Middleware.
func (m *middleware) Panic(c *gin.Context) {
	panic("unimplemented")
}

func (m *middleware) Authorize(c *gin.Context) {
	if strings.HasPrefix(c.Request.URL.Path, "/auth") {
		c.Next()
		return
	}

	header := c.GetHeader(values.Authorization)
	if header == "" {
		// _ = c.Error(values.ErrUnauthorized)
		c.Abort()
		return
	}

	token := strings.Split(header, "Bearer ")
	if len(token) != 2 { //nolint:gomnd
		// _ = c.Error(values.ErrUnauthorized)
		c.Abort()
		return
	}

	claims, err := m.jwt.ParseJWTToken(token[1])
	if err != nil {
		// _ = c.Error(values.NewHTTPError(http.StatusForbidden, err.Error()))
		c.Abort()
		return
	}

	c.Set("userID", claims.Subject)
	c.Next()

}

func New(jwt jwt.JWT) Middleware {
	return &middleware{
		jwt: jwt,
	}
}
