package utils

import (
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/mello369/loadbalancer/models"
)

func NewSimpleServer(addr string) *SimpleServer {
	serverUrl, err := url.Parse(addr)
	HandleError(err)
	return &models.SimpleServer{
		address: addr,
		proxy:   httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

func HandleError(err error) {
	if err != nil {
		os.Exit(1)
	}
}
