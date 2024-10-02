package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
	"github.com/naufalkhairil/golang-oauth-keycloak/config"
	"golang.org/x/oauth2"
)

type authToken struct {
	OAuth2Token *oauth2.Token
	// IDTokenClaims *json.RawMessage // ID Token payload is just JSON.
	IDTokenClaims map[string]interface{} // ID Token payload is just JSON.
}

type procTokenParams struct {
	username string
	email    string
	token    map[string]interface{}
}

var oauth2Config *oauth2.Config
var oidcConfig *oidc.Config
var provider *oidc.Provider

var procTokenStore map[string]*procTokenParams
var state string = "randomstate"

func InitMiddleware() {

	procTokenStore = make(map[string]*procTokenParams)

	ctx := context.Background()
	newProvider, err := oidc.NewProvider(ctx, config.GetProviderURL())
	if err != nil {
		log.Fatal(err)
	}
	provider = newProvider

	oauth2Config = &oauth2.Config{
		ClientID:     config.GetIDMID(),
		ClientSecret: config.GetIDMSecret(),
		RedirectURL:  config.GetCallbackURL(),
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	oidcConfig = &oidc.Config{
		ClientID: config.GetIDMID(),
	}
}

func Home(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("Waiting for authenticate\n\n%s \n", config.GetAuthURL()))
}

func Auth(c *gin.Context) {
	c.Redirect(http.StatusFound, oauth2Config.AuthCodeURL(state))
}

func AuthCallback(c *gin.Context) {
	if c.Request.URL.Query().Get("state") != state {
		c.JSON(http.StatusBadRequest, "state did not match")
		return
	}

	oauth2Token, err := oauth2Config.Exchange(c, c.Request.URL.Query().Get("code"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to exchange token: "+err.Error())
	}

	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, "No id_token field in oatuh2 token")
	}

	verifier := provider.Verifier(oidcConfig)
	idToken, err := verifier.Verify(c, rawIDToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to verify ID Token: "+err.Error())
	}

	tokenClaims := make(map[string]interface{})
	tokenClaims["oauth2token"] = oauth2Token
	tokenClaims["rawIDtoken"] = rawIDToken

	if err := idToken.Claims(&tokenClaims); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	email := tokenClaims["email"].(string)
	username := tokenClaims["preferred_username"].(string)

	procTokenStore[email] = &procTokenParams{
		username: username,
		email:    email,
		token:    tokenClaims,
	}

	c.Redirect(http.StatusTemporaryRedirect, "/success")
}

func GetToken(c *gin.Context) {
	email := c.Param("email")

	token, ok := procTokenStore[email]
	if !ok {
		c.JSON(http.StatusNotFound, map[string]string{
			"message": fmt.Sprintf("email %s not found", email),
		})
		return
	}

	c.JSON(http.StatusOK, token.token)
}

func Success(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(fmt.Sprintf(`
		<div style="
			background-color: #fff;
			padding: 40px;
			border-radius: 8px;
			box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
			text-align: center;
		">
			<h1 style="
				color: #333;
				margin-bottom: 20px;
			">You have Successfully signed in!</h1>
			</div>
		</div>
	`)))
}
