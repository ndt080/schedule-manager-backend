package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type TokenManager interface {
	NewAccessToken(userId int64) (string, error)
	NewRefreshToken(userId int64) (string, error)
	NewVerificationToken(userId int64, ttl time.Duration) (string, error)
	NewInviteToken(wid int64, ttl time.Duration) (string, error)
	ParseToken(tokenString string) (*TokenClaims, error)
	CheckPasswordHash(password, hash string) bool
}

type Manager struct {
	signingKey      string
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewManager(signingKey string, accessTokenTTL time.Duration, refreshTokenTTL time.Duration) (*Manager, error) {
	if signingKey == "" {
		return nil, errors.New("empty signing key")
	}
	return &Manager{
		signingKey:      signingKey,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
	}, nil
}

func (manager *Manager) NewAccessToken(userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.accessTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId:    userId,
		TokenType: "jwt",
	})

	return token.SignedString([]byte(manager.signingKey))
}

func (manager *Manager) NewRefreshToken(userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.refreshTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId:    userId,
		TokenType: "refresh",
	})

	return token.SignedString([]byte(manager.signingKey))
}

func (manager *Manager) NewVerificationToken(userId int64, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId:    userId,
		TokenType: "verification",
	})

	return token.SignedString([]byte(manager.signingKey))
}

func (manager *Manager) NewInviteToken(wid int64, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &InviteTokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		WorkspaceId: wid,
		TokenType:   "invite",
	})

	return token.SignedString([]byte(manager.signingKey))
}

func (manager *Manager) ParseToken(tokenString string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(manager.signingKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return nil, errors.New("token claims are not of type *tokenClaims")
	}

	return claims, err
}

func (manager *Manager) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
