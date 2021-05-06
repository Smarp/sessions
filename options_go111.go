// +build go1.11

package sessions

import "net/http"

// Options stores configuration for a session or session store.
//
// Fields are a subset of http.Cookie fields.
type Options struct {
	Path   string `json:"path"`
	Domain string `json:"domain"`
	// MaxAge=0 means no Max-Age attribute specified and the cookie will be
	// deleted after the browser session ends.
	// MaxAge<0 means delete cookie immediately.
	// MaxAge>0 means Max-Age attribute present and given in seconds.
	MaxAge   int `json:"originalMaxAge"`
	Secure   bool `json:"secure"`
	HttpOnly bool `json:"httpOnly"`
	// Defaults to http.SameSiteDefaultMode
	SameSite http.SameSite `json:"sameSite,omitempty"`
}
