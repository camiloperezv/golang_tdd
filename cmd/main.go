package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/camiloperezv/samples/users_db/pkg/dbuser"
	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hi there, I lovess %s!", r)
}

func declareRoutes(router *httprouter.Router) {
	router.GET("/", index)
	router.POST("/user", dbuser.NewUser)
}

func main() {
	router := httprouter.New()
	declareRoutes(router)
	log.Fatal(http.ListenAndServe(":3001", router))
}
