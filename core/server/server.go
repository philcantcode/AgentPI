package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/philcantcode/AgentPI/core"
)

var router *mux.Router
var rootLocation = "/"
var serverPort = "9100"

func server() {
	router = mux.NewRouter()

	fileServer := http.FileServer(http.Dir(rootLocation))
	router.PathPrefix("/").Handler(http.StripPrefix("/", fileServer))

	err := http.ListenAndServe(":"+serverPort, router)
	core.Error("Couldn't start server", err)
}
