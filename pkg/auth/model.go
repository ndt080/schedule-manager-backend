package auth

import "github.com/dgrijalva/jwt-go"

type TokenClaims struct {
	jwt.StandardClaims
	UserId    int64  `json:"userId"`
	TokenType string `json:"tokenType"`
}

type InviteTokenClaims struct {
	jwt.StandardClaims
	WorkspaceId int64  `json:"workspaceId"`
	TokenType   string `json:"tokenType"`
}
