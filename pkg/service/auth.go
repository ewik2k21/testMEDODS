package service

import (
	"crypto/sha512"
	"errors"
	"fmt"
	"math/rand"
	testmedods "testMEDODS"
	"testMEDODS/pkg/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user testmedods.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

const (
	salt       = "aksg838tkdnsg"
	signingKey = ("grksdjaj35j3a5j^J#$%3asdja")
	tokenttl   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int    `json:"user_id"`
	UserIP string `json:"user_ip"`
}

func generatePasswordHash(password string) string {
	hash := sha512.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	user, err := s.repo.GetUser(email, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, nil
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims not type")
	}

	return claims.UserId, nil
}

func (s *AuthService) NewRefreshToken() (string, error) {
	b := make([]byte, 32)

	n := rand.NewSource(time.Now().Unix())
	r := rand.New(n)

	_, err := r.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}
