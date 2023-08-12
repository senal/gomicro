package main

import (
	"fmt"
	"log"
	"net/http"
)

// why we need github.com/go-chi/chi/v5 <-- it's a 3rd party routing package
// any alternative for chi

//----- all the dependent packages
/*
	github.com/go-chi/chi/v5
	github.com/go-chi/chi/v5/middleware
	github.com/go-chi/cors
*/

const webPort = "80"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting broker service on port %s\n", webPort)

	// define http server
	//----> why refer http with an '&' <--- reference
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
