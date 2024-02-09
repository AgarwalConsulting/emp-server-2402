package main

import (
	"log"
	"net/http"

	empHTTP "algogrit.com/emp-server/employees/http"
	"algogrit.com/emp-server/employees/repository"
	"algogrit.com/emp-server/employees/service"
	"algogrit.com/emp-server/pkg/middleware"
)

func main() {
	var empRepo = repository.NewInMem()
	// config := Config{}
	// var empRepo, err = repository.NewSQL(config)

	// if err != nil {
	// 	log.Fatal("Unable to connect to DB:", err)
	// }

	var empV1Svc = service.NewV1(empRepo)
	var empHandler = empHTTP.NewHandler(empV1Svc)

	log.Println("Starting server on port: 8000...")
	http.ListenAndServe(":8000", middleware.Logging(empHandler))

	// http.ListenAndServe(":8000", handlers.CombinedLoggingHandler(os.Stdout, empHandler))
}
