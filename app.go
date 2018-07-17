package main

import (
	"net/http"
)

func main() {
	port := "3100"
	// w := utils.Logger.Writer()

	r := newRouter()
	// loggedRouter := handlers.LoggingHandler(w, r)
	http.ListenAndServe(":"+port, r)

}
