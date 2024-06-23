package types

import "github.com/golang-jwt/jwt/v5"

type ZitadelClaims struct {
	Roles map[string]map[string]string `json:"urn:zitadel:iam:org:project:roles"`
	jwt.RegisteredClaims
}
