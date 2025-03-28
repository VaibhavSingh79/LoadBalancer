package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type Server interface {
	Address() string
	IsAlive() bool
	Server(rw http.ResponseWriter, r *http.Request)
}

// this function will return the server
type SimplServer struct {
	addr string
	// ReverseProxy is an HTTP Handler that takes an incoming request and sends it to another server
	// proxying the response back to the client.
	proxy *httputil.ReverseProxy
}

func NewSimpleServer(addr string) *SimplServer {
	serverUrl, err := url.Parse(addr)
	handlError(err)

	//this function will return new instance of a server
	return &SimplServer{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

type LoadBalancer struct {
	port          string
	roundRobinCnt int
	servers       []Server
}

func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:          port,
		roundRobinCnt: 0,
		servers:       servers,
	}
}

// here I am handling the error requests
func handlError(err error) {
	if err != nil {
		fmt.Printf("Error %v\n", err)
		os.Exit(1)
	}
}

func (s *SimplServer) Address() string { return s.addr }

func (s *SimplServer) IsAlive() bool { return true }

func (s *SimplServer) Server(rw http.ResponseWriter, req *http.Request) {
	s.proxy.ServeHTTP(rw, req)
}

func (lb *LoadBalancer) getNextAliveServer() Server {
	server := lb.servers[lb.roundRobinCnt%len(lb.servers)]
	for !server.IsAlive() {
		lb.roundRobinCnt++
		server = lb.servers[lb.roundRobinCnt%len(lb.servers)]
	}
	lb.roundRobinCnt++
	return server

}

func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, req *http.Request) {
	targetServer := lb.getNextAliveServer()
	fmt.Printf("Forwarding Request to address %s\n", targetServer.Address())
	targetServer.Server(rw, req)
}

func main() {
	servers := []Server{
		NewSimpleServer("https://www.facebook.com"),
		NewSimpleServer("https://www.bing.com"),
		NewSimpleServer("https://www.duckduckgo.com"),
	}
	lb := NewLoadBalancer("8000", servers)
	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.serveProxy(rw, req)
	}
	http.HandleFunc("/", handleRedirect)

	fmt.Printf("Serving requests at 'localhost:%s'\n", lb.port)
	http.ListenAndServe(":"+lb.port, nil)
}
