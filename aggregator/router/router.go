package router

import (
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
)

func GetRouter(clientsServiceURL string, destinationsServiceURL string, sourcesServiceURL string) *mux.Router {
	clientsServiceU, _ := url.Parse(clientsServiceURL)
	destinationsServiceU, _ := url.Parse(destinationsServiceURL)
	sourcesServiceU, _ := url.Parse(sourcesServiceURL)

	router := mux.NewRouter()
	router.Handle("/clients", httputil.NewSingleHostReverseProxy(clientsServiceU))
	router.Handle("/destinations", httputil.NewSingleHostReverseProxy(destinationsServiceU))
	router.Handle("/sources", httputil.NewSingleHostReverseProxy(sourcesServiceU))

	return router
}
