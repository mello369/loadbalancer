package utils

import (
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/mello369/loadbalancer/models"
)

func NewSimpleServer(addr string) *models.SimpleServer {
	serverUrl, err := url.Parse(addr)
	HandleError(err)
	return &models.SimpleServer{
		Address: addr,
		Proxy:   httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

func NewLoadBalancer(port string, servers []models.Server) *models.LoadBalancer {
	return &models.LoadBalancer{
		Port:            port,
		Servers:         servers,
		RoundRobinCount: 0,
	}

}

func HandleError(err error) {
	if err != nil {
		os.Exit(1)
	}
}
