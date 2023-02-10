package utils

import (
	"net/http"
	"time"
	"tracka/pkg/config"
)

type CookieClaim struct {
	Email, IP, UserAgent string
}

func (cookie *CookieClaim) ToMap() map[string]interface{} {
	claims := make(map[string]interface{})
	claims["email"] = cookie.Email
	claims["ip"] = cookie.IP
	claims["user-agent"] = cookie.UserAgent

	return claims
}

func GenerateCookie(claims *CookieClaim) (*http.Cookie, error) {
	cookieGen := config.Get().Cookie
	if claims == nil {
		cookie := http.Cookie{
			Name:     cookieGen.Name,
			Expires:  time.Now().Add(-(time.Hour * 24 * 30 * 12)),
			SameSite: http.SameSiteNoneMode,
			Secure:   false,
			Path:     "/",
			Value:    "none",
		}
		return &cookie, nil
	}
	value, err := cookieGen.Cookie.Encode("logistics", claims.ToMap())
	if err != nil {
		return nil, err
	}

	cookie := http.Cookie{
		Name:     cookieGen.Name,
		Expires:  time.Now().Add(time.Hour * 24),
		SameSite: http.SameSiteNoneMode,
		Secure:   false,
		Value:    value,
		Path:     "/",
	}

	return &cookie, nil
}
