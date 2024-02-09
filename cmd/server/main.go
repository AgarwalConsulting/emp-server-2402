package main

import (
	"log"
	"net/http"
	"time"

	empHTTP "algogrit.com/emp-server/employees/http"
	"algogrit.com/emp-server/employees/repository"
	"algogrit.com/emp-server/employees/service"
)

func loggingMiddleware(next http.Handler) http.Handler {
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

func main() {
	var empRepo = repository.NewInMem()
	var empV1Svc = service.NewV1(empRepo)
	var empHandler = empHTTP.NewHandler(empV1Svc)

	log.Println("Starting server on port: 8000...")
	http.ListenAndServe(":8000", loggingMiddleware(empHandler))
	// http.ListenAndServe(":8000", handlers.CombinedLoggingHandler(os.Stdout, empHandler))
}
