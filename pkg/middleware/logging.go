package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	router := func(w http.ResponseWriter, req *http.Request) {
		begin := time.Now()

		// if req.Header // if authenticated!
		next.ServeHTTP(w, req)
		// else
		// w.WriteHeader(http.StatusUnauthorized)

		log.Printf("%s %s took %s\n", req.Method, req.URL, time.Since(begin))
	}

	return http.HandlerFunc(router)
}
