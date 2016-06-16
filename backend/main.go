package main

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"fmt"
	"net/http"
)

func main() {
	// Create our light-weight web serverish net/http handler
	n := negroni.Classic()

	// Create your servers router
	router := mux.NewRouter()

	//router.HandleFunc("/", )

	n.UseHandler(router)

	fmt.Println("Listening on localhost:3000")
	http.ListenAndServe(":3000", n)
}

