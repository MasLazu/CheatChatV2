package helper

import (
	"net/http"
	"time"
)

func SetCookies(writer http.ResponseWriter, name string, value string, expires time.Time) {
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Expires:  expires,
		HttpOnly: true,
		Domain:   ".localhost",
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Secure:   false,
	}
	http.SetCookie(writer, &cookie)
}
