package handlers

import (
	"encoding/json"
	"github.com/NewChakrit/golang_web_development/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/http"
	"time"
)

var googleOauth2Config *oauth2.Config
var oauthStateString = "golang-web-development"

func init() {
	googleOauth2Config = &oauth2.Config{
		ClientID:     config.Config.GoogleClientID,
		ClientSecret: config.Config.GoogleClientSecret,
		RedirectURL:  config.Config.GoogleRedirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}

}

func HandleGoogleLogin(ctx *gin.Context) {
	url := googleOauth2Config.AuthCodeURL("golang-web-development", oauth2.AccessTypeOffline)
	ctx.Redirect(http.StatusFound, url)
}

func HandleGoogleCallback(ctx *gin.Context) {
	// validate the state
	state := ctx.Query("state")
	if state != oauthStateString {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid oAuth state parameter"})
		return
	}
	// validate the code and get userToken back
	code := ctx.Query("code")
	token, err := googleOauth2Config.Exchange(ctx, code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange code for token"})
	}

	// get user info using userToken
	client := googleOauth2Config.Client(ctx, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}

	defer resp.Body.Close()

	// send userInfo to frontend
	var userInfo struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Picture string `json:"picture"`
	}

	if err = json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode user info"})
		return
	}

	jwtToken, err := generateJWT(userInfo.Email, userInfo.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token":   jwtToken,
		"email":   userInfo.Email,
		"name":    userInfo.Name,
		"picture": userInfo.Picture,
	})

}

func generateJWT(email, name string) (string, error) {
	tokenInfo := jwt.MapClaims{
		"email": email,
		"name":  name,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // 1 day expiration
		"iat":   time.Now().Unix(),                     // issued at
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenInfo)

	return token.SignedString([]byte(config.Config.JwtSaltKey))
}
