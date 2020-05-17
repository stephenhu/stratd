package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	conf = flag.String("conf", APP_CONFIG, "configuration file")
)


func addr() string {
  return fmt.Sprintf(":%s", app.Server.Port)
} // addr


func initRouter() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/version", versionHandler)
	//router.HandleFunc("/api/games/{id:[0-9]+}", gameHandler)

	return router

} // initRouter


func main() {

	flag.Parse()

	parseConfig()

	log.Printf("%s v%s starting on port %s...\n", APP_NAME, APP_VERSION, app.Server.Port)

	router := initRouter()

	log.Fatal(http.ListenAndServe(addr(), router))

} // main
