package auth

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/lukeshay/pms/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type SigningMethod string

const (
	SigningMethodHMAC   SigningMethod = "HMAC"
	SigningMethodECDSA  SigningMethod = "ECDSA"
	SigningMethodRSA    SigningMethod = "RSA"
	SigningMethodRSAPSS SigningMethod = "RSAPSS"
)

type Claims struct {
	jwt.RegisteredClaims
	models.User
}

type Auth struct {
	JWTSecret     string
	FindSecret    func(claims *Claims) ([]byte, error)
	SigningMethod SigningMethod
}

func (a *Auth) PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (a *Auth) PasswordHashCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (a *Auth) verifySigningMethod(token *jwt.Token) bool {
	if a.SigningMethod == SigningMethodHMAC {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		return ok
	} else if a.SigningMethod == SigningMethodRSAPSS {
		_, ok := token.Method.(*jwt.SigningMethodRSAPSS)

		return ok
	} else if a.SigningMethod == SigningMethodECDSA {
		_, ok := token.Method.(*jwt.SigningMethodECDSA)

		return ok
	} else if a.SigningMethod == SigningMethodRSA {
		_, ok := token.Method.(*jwt.SigningMethodRSA)

		return ok
	}

	return false
}

func (a *Auth) JWTParse(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if ok := a.verifySigningMethod(token); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		if a.JWTSecret != "" {
			return []byte(a.JWTSecret), nil
		}

		claims, ok := token.Claims.(*Claims)
		if !ok {
			return nil, fmt.Errorf("invalid claims")
		}

		return a.FindSecret(claims)
	})
	if err != nil {
		return nil, err
	}

	fmt.Printf("claims: %+v\n", token.Claims)

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}

	return claims, nil

}

func (a *Auth) getSigningMethod() jwt.SigningMethod {
	if a.SigningMethod == SigningMethodHMAC {
		return jwt.SigningMethodHS256
	} else if a.SigningMethod == SigningMethodRSAPSS {
		return jwt.SigningMethodPS256
	} else if a.SigningMethod == SigningMethodECDSA {
		return jwt.SigningMethodES256
	} else if a.SigningMethod == SigningMethodRSA {
		return jwt.SigningMethodRS256
	}

	return nil
}

func (a *Auth) JWTGenerate(user *models.User) (string, error) {
	t := time.Now()

	claims := &Claims{
		jwt.RegisteredClaims{
			Subject:   user.Id,
			Issuer:    "pms",
			Audience:  jwt.ClaimStrings{"pms"},
			ExpiresAt: jwt.NewNumericDate(t.AddDate(0, 0, 7)),
			NotBefore: jwt.NewNumericDate(t),
			IssuedAt:  jwt.NewNumericDate(t),
			ID:        uuid.NewString(),
		},
		*user,
	}
	claims.Password = ""

	token := jwt.NewWithClaims(a.getSigningMethod(), claims)

	if a.JWTSecret != "" {
		return token.SignedString([]byte(a.JWTSecret))
	}

	secret, err := a.FindSecret(claims)
	if err != nil {
		return "", err
	}

	return token.SignedString(secret)
}

func SetClaims(ctx *gin.Context, claims *Claims) {
	ctx.Set("claims", claims)
}

func GetClaims(ctx *gin.Context) (*Claims, error) {
	claims, present := ctx.Get("claims")
	if !present {
		return nil, fmt.Errorf("claims not present in context")
	}

	authClaims, ok := claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("claims not of type auth.Claims")
	}

	return authClaims, nil
}
func RequireClaims(ctx *gin.Context) *Claims {
	claims, _ := GetClaims(ctx)

	return claims
}
