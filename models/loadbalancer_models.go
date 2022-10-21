package models

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

type Server interface {
	getAddress() string
	isAlive() bool
	Serve(rw http.ResponseWriter, rq *http.Request)
}

type SimpleServer struct {
	Address string
	Proxy   *httputil.ReverseProxy
}

func (s *SimpleServer) getAddress() string {
	return s.Address
}

func (s *SimpleServer) isAlive() bool {
	return true
}

func (s *SimpleServer) Serve(rw http.ResponseWriter, rq *http.Request) {
	s.Proxy.ServeHTTP(rw, rq)
}

type LoadBalancer struct {
	Servers         []Server
	Port            string
	RoundRobinCount int
}

func (lb *LoadBalancer) ServeProxy(rw http.ResponseWriter, rq *http.Request) {
	targetServer := lb.getNextAvailableServer()
	fmt.Printf("Forwarding requesting to address : %v", targetServer.getAddress())
	targetServer.Serve(rw, rq)
}

func (lb *LoadBalancer) getNextAvailableServer() Server {
	server := lb.Servers[lb.RoundRobinCount%len(lb.Servers)]
	for !server.isAlive() {
		lb.RoundRobinCount++
		server = lb.Servers[lb.RoundRobinCount%len(lb.Servers)]
	}
	lb.RoundRobinCount++
	return server
}
