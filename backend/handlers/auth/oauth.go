package auth

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/google/uuid"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

var (
	oauthConfig *oauth2.Config
	verifier    *oidc.IDTokenVerifier
	state       = uuid.New().String()
)

func InitializeOauth() {
	clientID := os.Getenv("CLIENT_ID")
	providerUrl := os.Getenv("PROVIDER_URL")
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, providerUrl)
	if err != nil {
		panic(err)
	}
	verifier = provider.Verifier(&oidc.Config{ClientID: clientID})

	oauthConfig = &oauth2.Config{
		ClientID:    clientID,
		RedirectURL: os.Getenv("REDIRECT_URL"),
		Scopes:      []string{"openid", "profile", "email", "offline_access"},
		Endpoint:    provider.Endpoint(),
	}
}

func LoginHandler(c echo.Context) error {
	verifier := oauth2.GenerateVerifier()
	authURL := oauthConfig.AuthCodeURL(state, oauth2.S256ChallengeOption(verifier))

	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	sess.Values["code_verifier"] = verifier

	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}

	return c.Redirect(http.StatusFound, authURL)
}

func CallbackHandler(c echo.Context) error {
	if c.QueryParam("state") != state {
		return c.String(http.StatusBadRequest, "Invalid state parameter")
	}

	code := c.QueryParam("code")
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	codeVerifier := sess.Values["code_verifier"].(string)

	token, err := oauthConfig.Exchange(c.Request().Context(), code, oauth2.VerifierOption(codeVerifier))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to exchange token")
	}

	rawIdToken := token.Extra("id_token").(string)
	slog.Info("raw token", "raw_token", rawIdToken)

	cookie := new(http.Cookie)
	cookie.Name = "access_token"
	cookie.Value = token.AccessToken
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	c.SetCookie(cookie)

	cookie = new(http.Cookie)
	cookie.Name = "id_token"
	cookie.Value = rawIdToken
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

func LoginCheckHandler(c echo.Context) error {
	rawIdToken, err := c.Cookie("id_token")
	idToken, err := verifier.Verify(c.Request().Context(), rawIdToken.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid id token")
	}
	var claims struct {
		GivenName  string `json:"given_name"`
		FamilyName string `json:"family_name"`
		Username   string `json:"preferred_username"`
		Email      string `json:"email"`
	}
	idToken.Claims(&claims)
	return c.JSON(http.StatusOK, claims)
}
