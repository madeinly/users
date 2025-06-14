package auth

import (
	"net/http"
	"time"
)

type Cookie struct {
	Name     string        `json:"name"`
	Value    string        `json:"value"`
	Path     string        `json:"path"`
	HttpOnly bool          `json:"httpOnly,omitempty"`
	Secure   bool          `json:"secure,omitempty"`
	SameSite http.SameSite `json:"sameSite"`
	Expires  time.Time     `json:"expires"`
}

const CookieName = "capre_token"

func SetCookie(userId string, email string, w http.ResponseWriter, r *http.Request) (Cookie, error) {
	token, err := GenerateToken(userId, email)
	if err != nil {
		return Cookie{}, err
	}

	expires := time.Now().Add(7 * 24 * time.Hour)

	cookie := Cookie{
		Name:  CookieName,
		Value: token,
		Path:  "/",
		// HttpOnly: true,
		// Secure:   true,
		SameSite: http.SameSiteNoneMode,
		Expires:  expires,
	}

	return cookie, nil
}

func ValidateCookie(r *http.Request) (bool, Claims, error) {

	// Get the cookie from the request
	cookie, err := r.Cookie(CookieName)

	if err != nil {
		return false, Claims{}, err
	}

	claims, err := ValidateToken(cookie.Value)

	if err != nil {
		return false, *claims, err
	}

	return true, *claims, nil

}
