package utils

import (
	"net/http"
	"strconv"
	"time"
)

type CookieClaim struct {
	Email, IP, UserAgent string
}

func GenerateCookie(claims CookieClaim) *http.Cookie {
	cookie := http.Cookie{
		Name:     "logistics",
		Expires:  time.Now().Add(time.Hour * 24),
		SameSite: http.SameSiteNoneMode,
		Secure:   false,
	}

	// placeholder
	cookie.Value = claims.Email + strconv.Itoa(int(cookie.Expires.Unix()))

	return &cookie
}
