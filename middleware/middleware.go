package middleware

import (
	"net/http"
	"strings"
)

func UseFormMethod(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			m := r.PostFormValue("_method")
			if m == "DELETE" || m == "PATCH" || m == "PUT" {
				r.Method = strings.ToUpper(m)
			}
		}

		next.ServeHTTP(w, r)
	})
}
