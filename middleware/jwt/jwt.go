package jwt

import (
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	errTokenExpired     = errors.New("token is expired")
	errTokenNotValidYet = errors.New("token not active yet")
	errTokenMalformed   = errors.New("that's not even a token")
	errTokenInvalid     = errors.New("couldn't handle this token")
	signKey             = "testSecretKey" // kind like a secret key
)

// JWT jwt instance
type JWT struct {
	SigningKey []byte
}

// CustomClaims custom claims
type CustomClaims struct {
	// UID      string `json:"uid"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	jwt.StandardClaims
}

// NewJWT make a new JWT, just a struct with a secret key string
func NewJWT() *JWT {
	return &JWT{[]byte(signKey)}
}

// CreateToken create a jwt token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	// encode algothrim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken parse a token string, return the claims or error
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errTokenNotValidYet
			} else {
				return nil, errTokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errTokenInvalid
}

// RefreshToken refresh a jwt
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", nil
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", errTokenInvalid
}

// Auth is a middleware for checking token
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 1,
				"msg":  "no authorization",
				"data": nil,
			})
			c.Abort()
			return
		}

		j := NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 1,
				"msg":  err.Error(),
				"data": nil,
			})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
