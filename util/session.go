package util

import "net/http"

func GetSessionID(r *http.Request) string {
	c, err := r.Cookie("PHPSESSID")
	if err != nil {
		return ""
	}

	return c.Value
}
