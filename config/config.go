package config

import (
	"time"

	"github.com/spf13/viper"
)

const (
	cfgIDMID              = "IDM_ID"
	cfgIDMSecret          = "IDM_SECRET"
	cfgMiddlewareBaseURL  = "MIDDLEWARE_BASE_URL"
	cfgMiddlewareAuthUrl  = "MIDDLEWARE_AUTH_URL"
	cfgMiddlewareTokenUrl = "MIDDLEWARE_TOKEN_URL"
	cfgMiddlewareCallback = "MIDDLEWARE_CALLBACK_URL"
	cfgAuthTimeout        = "CLIENT_AUTH_TIMEOUT"
	cfgProviderURL        = "PROVIDER_URL"
	cfgJWTSignKey         = "JWT_SIGNATURE_KEY"
	cfgJWTExpiredDuration = "JWT_EXPIRED_DURATION"
)

func GetBaseURL() string {
	return viper.GetString(cfgMiddlewareBaseURL)
}

func GetAuthURL() string {
	return viper.GetString(cfgMiddlewareAuthUrl)
}

func GetAuthTokenURL(email string) string {
	return viper.GetString(cfgMiddlewareTokenUrl) + "/" + email
}

func GetAuthTimeout() time.Duration {
	return viper.GetDuration(cfgAuthTimeout)
}

func GetIDMID() string {
	return viper.GetString(cfgIDMID)
}

func GetIDMSecret() string {
	return viper.GetString(cfgIDMSecret)
}

func GetCallbackURL() string {
	return viper.GetString(cfgMiddlewareCallback)
}

func GetProviderURL() string {
	return viper.GetString(cfgProviderURL)
}

func GetJWTSignKey() string {
	return viper.GetString(cfgJWTSignKey)
}

func GetJWTExpiredDuration() time.Duration {
	if viper.GetString(cfgJWTExpiredDuration) == "" {
		return time.Duration(1440 * time.Second)
	}

	return viper.GetDuration(cfgJWTExpiredDuration)
}
