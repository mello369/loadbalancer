package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mello369/loadbalancer/models"
	"github.com/mello369/loadbalancer/routes"
	"github.com/mello369/loadbalancer/utils"
)

func main() {
	servers := []models.Server{
		utils.NewSimpleServer("https://www.facebook.com"),
		utils.NewSimpleServer("https://www.google.com"),
		utils.NewSimpleServer("https://www.duckduckgo.com"),
	}
	lb := utils.NewLoadBalancer("8000", servers)
	routes.RegisterRoutes(lb)
	fmt.Printf("serving requests at port %v", lb.Port)
	log.Fatal(http.ListenAndServe(":"+lb.Port, nil))
}
