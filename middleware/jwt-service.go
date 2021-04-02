package middleware

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

var (
	issuer, secretKey string
	expires           int
)

func init() {
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("config.json")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
	issuer = viper.GetString("issuer")
	secretKey = viper.GetString("secretKey")

}

type JWTService interface {
	GenerateToken(userId string) string
	ValidateToken(token string) (*jwt.Token, error)
	Test(string) string
}

func (j *jwtService) Test(s string) string {
	return s

}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJwtService() JWTService {
	return &jwtService{
		issuer:    issuer,
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := secretKey
	return secretKey
}

func (j *jwtService) GenerateToken(userid string) string {
	claims := &jwtCustomClaim{
		userid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return ""
	}
	return tokenString
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("error signing method")
		}
		return []byte(j.secretKey), nil
	})
}
