package main

import (
	"net/http"

	router "github.com/user/go-test-rest/delivery"
)

func main() {
	port := "3100"
	// w := utils.Logger.Writer()

	r := router.UserRouter()
	// loggedRouter := handlers.LoggingHandler(w, r)
	http.ListenAndServe(":"+port, r)

}
