package handlers

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	customErrors "github.com/cuonglv-smartosc/golang-boiler-template/internal/applications/api/errors"
	"github.com/cuonglv-smartosc/golang-boiler-template/internal/applications/api/modules/auth/dtos"

	"github.com/cuonglv-smartosc/golang-boiler-template/internal/repository"
	repository_models "github.com/cuonglv-smartosc/golang-boiler-template/internal/repository/models"
	"github.com/cuonglv-smartosc/golang-boiler-template/pkg/auth"
)

type AuthService struct {
	db         repository.Storage
	jwtService *auth.JWTService
}

func NewAuthService(db repository.Storage) *AuthService {
	return &AuthService{
		db:         db,
		jwtService: auth.NewJWTService(),
	}
}

func (s *AuthService) Login(req *dtos.LoginRequest) (*dtos.LoginResponse, error) {
	user, err := s.db.GetUserByEmail(context.Background(), req.Email)
	if err != nil {
		return nil, customErrors.NewNotFoundError("User", req.Email)
	}

	//if !auth.CheckPasswordHash(req.Password, user.Password) {
	//	return nil, customErrors.NewAuthError("INVALID_PASSWORD", "Invalid password provided")
	//}

	tokenPair, err := s.jwtService.GenerateTokenPair(
		strconv.FormatInt(user.ID, 10),
		user.Email,
		[]string{"user"},
	)
	if err != nil {
		return nil, customErrors.NewAuthError("TOKEN_GENERATION_FAILED", "Failed to generate authentication tokens")
	}

	return &dtos.LoginResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
		ExpiresAt:    tokenPair.ExpiresAt.Unix(),
		User: dtos.UserInfo{
			ID:    user.ID,
			Email: user.Email,
			Roles: []string{"user"},
		},
	}, nil
}

func (s *AuthService) RefreshToken(req *dtos.RefreshTokenRequest) (*dtos.RefreshTokenResponse, error) {
	tokenPair, err := s.jwtService.RefreshAccessToken(req.RefreshToken)
	if err != nil {
		return nil, customErrors.NewAuthError("INVALID_REFRESH_TOKEN", "Refresh token is invalid or expired")
	}

	return &dtos.RefreshTokenResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
		ExpiresAt:    tokenPair.ExpiresAt.Unix(),
	}, nil
}

func (s *AuthService) GetUserByID(id int64) (*dtos.UserInfo, error) {
	user, err := s.db.GetUserByID(context.Background(), id)
	if err != nil {
		return nil, customErrors.NewNotFoundError("User", strconv.FormatInt(id, 10))
	}

	return &dtos.UserInfo{
		ID:    user.ID,
		Email: user.Email,
		Roles: []string{"user"},
	}, nil
}

func (s *AuthService) Register(req *dtos.RegisterRequest) (*dtos.RegisterResponse, error) {
	if _, err := s.db.GetUserByEmail(context.Background(), req.Email); err == nil {
		return nil, customErrors.NewCustomError(http.StatusBadRequest, "EMAIL_EXISTS", "Email already registered", nil)
	}

	hashed, err := auth.HashPassword(req.Password)
	if err != nil {
		return nil, customErrors.NewAuthError("HASH_FAILED", "Failed to hash password")
	}

	user := &repository_models.User{
		Email:       strings.ToLower(req.Email),
		PhoneNumber: req.PhoneNumber,
		Password:    hashed,
	}

	if err := s.db.CreateUser(user); err != nil {
		return nil, customErrors.NewBusinessLogicError("CREATE_USER_FAILED", "Failed to create user")
	}

	return &dtos.RegisterResponse{
		ID:    user.ID,
		Email: user.Email,
		Roles: []string{"user"},
	}, nil
}
