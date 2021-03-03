package shared

import (
	"net/http"
)

var settings = struct {
	DisableSecure bool
}{
	DisableSecure: false,
}

func Init(disableSecure bool) {
	settings.DisableSecure = disableSecure
}

func GetCookie(req *http.Request, name string) (string, bool) {
	cookie, err := req.Cookie(name)
	if err != nil {
		return ``, false
	}
	return cookie.Value, true
}

func SetCookie(w http.ResponseWriter, name, value string, maxAgeInSeconds int) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    value,
		HttpOnly: true,
		Secure:   !settings.DisableSecure,
		MaxAge:   maxAgeInSeconds,
	})
}

func UnsetCookie(w http.ResponseWriter, name string) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    ``,
		HttpOnly: true,
		Secure:   !settings.DisableSecure,
		MaxAge:   -1,
	})
}
