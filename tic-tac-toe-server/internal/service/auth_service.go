package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/common"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/config"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repository.UserRepository
}

func NewAuthService(repoRoom repository.UserRepository) *AuthService {
	return &AuthService{
		repo: repoRoom,
	}
}

type Claims struct {
	Sub struct {
		Email string `json:"email"`
	} `json:"sub"`
	jwt.RegisteredClaims
}

func CheckTokenIsNotExpired(token string) (*Claims, error) {
	token = strings.TrimSpace(strings.ReplaceAll(token, "Bearer ", ""))

	claims, err := parseToken(token)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func (service *AuthService) SignIn(ctx context.Context, form common.AuthSignInRequest) (map[string]string, error) {
	currentUser, err := service.repo.FindByEmail(ctx, form.Email)
	if err != nil {
		return nil, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(currentUser.Password), []byte(strings.TrimSpace(form.Password))); err != nil {
		return nil, errors.New("password is not valid")
	}

	accessToken, err := getToken(*currentUser)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"token": accessToken,
	}, nil
}

func (service *AuthService) SignUp(ctx context.Context, form common.AuthSignUpRequest) error {
	if _, err := service.repo.FindByEmail(ctx, form.Email); err == nil {
		return errors.New("user with this email already exists")
	}
	password, err := bcrypt.GenerateFromPassword(
		[]byte(strings.TrimSpace(form.Password)),
		config.ServerConfig.BcryptPower,
	)
	if err != nil {
		return err
	}
	form.Password = string(password)
	return service.repo.Create(ctx, form)
}

func getToken(user common.User) (string, error) {
	seconds := config.ServerConfig.JWTConfig.AccessTokenTTL
	duration, err := time.ParseDuration(seconds)

	if err != nil {
		return "", err
	}
	payload := jwt.MapClaims{
		"sub": map[string]interface{}{
			"email": user.Email,
		},
		"exp": time.Now().Add(time.Duration(duration)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	jwtSecret := []byte(config.ServerConfig.JWTConfig.AccessTokenSecret)
	t, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return t, nil
}

func parseToken(token string) (*Claims, error) {
	var claims Claims
	t, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.ServerConfig.JWTConfig.AccessTokenSecret), nil
	})

	if claims.ExpiresAt != nil && time.Now().Unix() > claims.ExpiresAt.Unix() {
		return nil, errors.New("token is expired")
	}

	if err != nil || !t.Valid {
		return nil, errors.New("your token is invalid")
	}

	return &claims, nil
}
