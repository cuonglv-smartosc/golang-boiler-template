package auth

import (
	"errors"
	"time"

	"github.com/cuonglv-smartosc/golang-boiler-template/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token has expired")
)

type JWTService struct {
	secretKey       string
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

type Claims struct {
	UserID    string   `json:"user_id"`
	Email     string   `json:"email"`
	Roles     []string `json:"roles"`
	TokenType string   `json:"token_type"` // "access" or "refresh"
	jwt.RegisteredClaims
}

type TokenPair struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

// NewJWTService tạo JWT service mới
func NewJWTService() *JWTService {
	return &JWTService{
		secretKey:       config.Default.Jwt.Secret,
		accessTokenTTL:  15 * time.Minute,   // Access token hết hạn sau 15 phút
		refreshTokenTTL: 7 * 24 * time.Hour, // Refresh token hết hạn sau 7 ngày
	}
}

// NewJWTServiceWithConfig tạo JWT service với config tùy chỉnh
func NewJWTServiceWithConfig(secretKey string, accessTTL, refreshTTL time.Duration) *JWTService {
	return &JWTService{
		secretKey:       secretKey,
		accessTokenTTL:  accessTTL,
		refreshTokenTTL: refreshTTL,
	}
}

// GenerateTokenPair tạo cả access token và refresh token
func (s *JWTService) GenerateTokenPair(userID, email string, roles []string) (*TokenPair, error) {
	// Tạo access token
	accessToken, err := s.generateToken(userID, email, roles, "access", s.accessTokenTTL)
	if err != nil {
		return nil, err
	}

	// Tạo refresh token
	refreshToken, err := s.generateToken(userID, email, roles, "refresh", s.refreshTokenTTL)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(s.accessTokenTTL),
	}, nil
}

// GenerateAccessToken tạo access token mới
func (s *JWTService) GenerateAccessToken(userID, email string, roles []string) (string, error) {
	return s.generateToken(userID, email, roles, "access", s.accessTokenTTL)
}

// GenerateRefreshToken tạo refresh token mới
func (s *JWTService) GenerateRefreshToken(userID, email string, roles []string) (string, error) {
	return s.generateToken(userID, email, roles, "refresh", s.refreshTokenTTL)
}

// generateToken tạo JWT token với thông tin đã cho
func (s *JWTService) generateToken(userID, email string, roles []string, tokenType string, ttl time.Duration) (string, error) {
	now := time.Now()
	expiresAt := now.Add(ttl)

	claims := Claims{
		UserID:    userID,
		Email:     email,
		Roles:     roles,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "your-app-name",
			Subject:   userID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken xác thực token và trả về claims
func (s *JWTService) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Kiểm tra signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.secretKey), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

// ValidateAccessToken xác thực access token
func (s *JWTService) ValidateAccessToken(tokenString string) (*Claims, error) {
	claims, err := s.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	if claims.TokenType != "access" {
		return nil, errors.New("token is not an access token")
	}

	return claims, nil
}

// ValidateRefreshToken xác thực refresh token
func (s *JWTService) ValidateRefreshToken(tokenString string) (*Claims, error) {
	claims, err := s.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	if claims.TokenType != "refresh" {
		return nil, errors.New("token is not a refresh token")
	}

	return claims, nil
}

// RefreshAccessToken tạo access token mới từ refresh token
func (s *JWTService) RefreshAccessToken(refreshToken string) (*TokenPair, error) {
	// Validate refresh token
	claims, err := s.ValidateRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}

	// Tạo token pair mới
	return s.GenerateTokenPair(claims.UserID, claims.Email, claims.Roles)
}

// GetUserIDFromToken lấy user ID từ token
func (s *JWTService) GetUserIDFromToken(tokenString string) (string, error) {
	claims, err := s.ValidateToken(tokenString)
	if err != nil {
		return "", err
	}
	return claims.UserID, nil
}

// GetRolesFromToken lấy roles từ token
func (s *JWTService) GetRolesFromToken(tokenString string) ([]string, error) {
	claims, err := s.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}
	return claims.Roles, nil
}

// IsTokenExpired kiểm tra token đã hết hạn chưa
func (s *JWTService) IsTokenExpired(tokenString string) bool {
	claims, err := s.ValidateToken(tokenString)
	if err != nil {
		return true
	}
	return claims.ExpiresAt.Before(time.Now())
}
