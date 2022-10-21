package routes

import (
	"net/http"

	"github.com/mello369/loadbalancer/models"
)

var RegisterRoutes = func(lb *models.LoadBalancer) {
	handleRedirect := func(rw http.ResponseWriter, rq *http.Request) {
		lb.ServeProxy(rw, rq)

	}
	http.HandleFunc("/", handleRedirect)
}
