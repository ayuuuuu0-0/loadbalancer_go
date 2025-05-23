// package main

// import (
// 	"fmt"
// 	"net/http/httputil"
// 	"net/http"
// 	"net/url"
// 	"os"
// )

// type Server interface {
// 	Address() string
// 	IsAlive () bool
// 	Serve(rw http.ResponseWriter, r *http.Request)
// }

// type simpleServer struct {
// 	addr string
// 	proxy *httputil.ReverseProxy
// }

// func newSimpleServer(addr string) *simpleServer {
// 	serverUrl, err := url.Parse(addr)
// 	handleErr(err)

// 	return &simpleServer{
// 		addr: addr,
// 		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
// 	}
// }

// type LoadBalancer struct {
// 	port               string
// 	roundRobinCount    int
// 	servers            []Server
// }

// func NewLoadBalancer(port string, servers []Server) *LoadBalancer{
// 	return &LoadBalancer{
// 		port: port,
// 		roundRobinCount: 0,
// 		servers: servers,
// 	}

// }

// func handleErr(err error) {
// 	if err != nil {
// 		fmt.Printf("error: %v\n", err)
// 		os.Exit(1)
// 	}
// }

// func (s * simpleServer) Address() string {return s.addr}

// func (s * simpleServer) IsAlive() bool {return true}

// func (s * simpleServer) Serve(rw http.ResponseWriter, req *http.Request) {
// 	s.proxy.ServeHTTP(rw, req)
// }

// func (lb * LoadBalancer) getNexAvailableServer() Server {
// 	server := lb.servers[lb.roundRobinCount%len(lb.servers)]
// 	for !server.IsAlive() {
// 		lb.roundRobinCount++
// 		server = lb.servers[lb.roundRobinCount%len(lb.servers)]
// 	}
// 	lb.roundRobinCount++
// 	return server
// }

// func (lb * LoadBalancer) serveProxy(rw http.ResponseWriter, req *http.Request){
// 	targetServer := lb.getNexAvailableServer()
// 	fmt.Printf("forwarding request to %s\n", targetServer.Address())
// 	targetServer.Serve(rw, req)
// }

// func main(){
// 	servers := []Server{
// 		newSimpleServer("http://instagram.com"),
// 		newSimpleServer("http://facebook.com"),
// 		newSimpleServer("http://google.com"),
// 	}
// 	lb := NewLoadBalancer("8000", servers)
// 	handleRedirect := func(rw http.ResponseWriter, req *http.Request){
// 		lb.serveProxy(rw, req)
// 	}
// 	http.HandleFunc("/" , handleRedirect)

// 	fmt.Printf("serving on port %s\n", lb.port)
//     http.ListenAndServe(":"+lb.port, nil)
// }

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"net/http/httputil"
// 	"net/url"
// 	"sync"
// )

// type Server interface {
// 	Address() string
// 	IsAlive() bool
// 	Serve(http.ResponseWriter, *http.Request)
// }

// type simpleServer struct {
// 	addr  string
// 	proxy *httputil.ReverseProxy
// }

// func newSimpleServer(addr string) *simpleServer {
// 	serverURL, err := url.Parse(addr)
// 	if err != nil {
// 		log.Fatalf("Error parsing URL: %v", err)
// 	}

// 	return &simpleServer{
// 		addr:  addr,
// 		proxy: httputil.NewSingleHostReverseProxy(serverURL),
// 	}
// }

// type LoadBalancer struct {
// 	port            string
// 	roundRobinCount int
// 	servers         []Server
// 	mu              sync.Mutex
// }

// func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
// 	return &LoadBalancer{
// 		port:            port,
// 		roundRobinCount: 0,
// 		servers:         servers,
// 	}
// }

// func (s *simpleServer) Address() string {
// 	return s.addr
// }

// func (s *simpleServer) IsAlive() bool {
// 	resp, err := http.Get(s.addr)
// 	if err != nil {
// 		return false
// 	}
// 	defer resp.Body.Close()
// 	return resp.StatusCode == http.StatusOK
// }

// func (s *simpleServer) Serve(rw http.ResponseWriter, req *http.Request) {
// 	s.proxy.ServeHTTP(rw, req)
// }

// func (lb *LoadBalancer) getNextAvailableServer() Server {
// 	lb.mu.Lock()
// 	defer lb.mu.Unlock()

// 	for i := 0; i < len(lb.servers); i++ {
// 		server := lb.servers[lb.roundRobinCount%len(lb.servers)]
// 		lb.roundRobinCount++
// 		if server.IsAlive() {
// 			return server
// 		}
// 	}
// 	log.Fatal("No available servers")
// 	return nil
// }

// func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, req *http.Request) {
// 	targetServer := lb.getNextAvailableServer()
// 	fmt.Printf("Forwarding request to %s\n", targetServer.Address())
// 	targetServer.Serve(rw, req)
// }

// func main() {
// 	servers := []Server{
// 		newSimpleServer("https://github.com"),
// 		newSimpleServer("https://www.duckduckgo.com"),
// 		newSimpleServer("https://college-cart.vercel.app"),
// 	}

// 	lb := NewLoadBalancer("8000", servers)

// 	http.HandleFunc("/", lb.serveProxy)

// 	fmt.Printf("Serving on localhost:%s\n", lb.port)
// 	log.Fatal(http.ListenAndServe("localhost:"+lb.port, nil))
// }

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"net/http/httputil"
// 	"net/url"
// 	"sync"
// )

// type Server interface {
// 	Address() string
// 	IsAlive() bool
// 	Serve(http.ResponseWriter, *http.Request)
// }

// type simpleServer struct {
// 	addr  string
// 	proxy *httputil.ReverseProxy
// }

// func newSimpleServer(addr string) *simpleServer {
// 	serverURL, err := url.Parse(addr)
// 	if err != nil {
// 		log.Fatalf("Error parsing URL: %v", err)
// 	}

// 	proxy := httputil.NewSingleHostReverseProxy(serverURL)
	
// 	// Modify request to preserve original host and IP
// 	proxy.Director = func(req *http.Request) {
// 		req.URL = serverURL
// 		req.Host = serverURL.Host
// 		req.Header.Set("X-Forwarded-Host", req.Host)
// 		req.Header.Set("X-Origin-Host", req.Host)
// 	}

// 	return &simpleServer{
// 		addr:  addr,
// 		proxy: proxy,
// 	}
// }

// type LoadBalancer struct {
// 	port            string
// 	roundRobinCount int
// 	servers         []Server
// 	mu              sync.Mutex
// }

// func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
// 	return &LoadBalancer{
// 		port:            port,
// 		roundRobinCount: 0,
// 		servers:         servers,
// 	}
// }

// func (s *simpleServer) Address() string {
// 	return s.addr
// }

// func (s *simpleServer) IsAlive() bool {
// 	client := &http.Client{
// 		CheckRedirect: func(req *http.Request, via []*http.Request) error {
// 			return http.ErrUseLastResponse
// 		},
// 	}
// 	resp, err := client.Get(s.addr)
// 	if err != nil {
// 		return false
// 	}
// 	defer resp.Body.Close()
// 	return resp.StatusCode == http.StatusOK || resp.StatusCode < 400
// }

// func (s *simpleServer) Serve(rw http.ResponseWriter, req *http.Request) {
// 	s.proxy.ServeHTTP(rw, req)
// }

// func (lb *LoadBalancer) getNextAvailableServer() Server {
// 	lb.mu.Lock()
// 	defer lb.mu.Unlock()

// 	for i := 0; i < len(lb.servers); i++ {
// 		server := lb.servers[lb.roundRobinCount%len(lb.servers)]
// 		lb.roundRobinCount++
// 		if server.IsAlive() {
// 			return server
// 		}
// 	}
// 	log.Fatal("No available servers")
// 	return nil
// }

// func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, req *http.Request) {
// 	targetServer := lb.getNextAvailableServer()
// 	fmt.Printf("Forwarding request to %s\n", targetServer.Address())
// 	targetServer.Serve(rw, req)
// }

// func main() {
// 	servers := []Server{
// 		newSimpleServer("https://www.instagram.com"),
// 		newSimpleServer("https://www.facebook.com"),
// 		newSimpleServer("https://college-cart.vercel.app"),
// 		// Add your personal project URL here
// 	}

// 	lb := NewLoadBalancer("8000", servers)

// 	http.HandleFunc("/", lb.serveProxy)

// 	fmt.Printf("Serving on localhost:%s\n", lb.port)
// 	log.Fatal(http.ListenAndServe("localhost:"+lb.port, nil))
// }

//FOURTH VERSION

package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"strings"
)

type Server interface {
	Address() string
	IsAlive() bool
	Serve(http.ResponseWriter, *http.Request)
}

type simpleServer struct {
	addr  string
	proxy *httputil.ReverseProxy
}

func newSimpleServer(addr string) *simpleServer {
	serverURL, err := url.Parse(addr)
	if err != nil {
		log.Fatalf("Error parsing URL: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(serverURL)
	
	proxy.Director = func(req *http.Request) {
		req.URL = serverURL
		req.Host = serverURL.Host
		
		// Specific handling for module scripts
		if strings.Contains(req.URL.Path, ".js") {
			req.Header.Set("Content-Type", "application/javascript")
		}
	}

	proxy.ModifyResponse = func(resp *http.Response) error {
		// Ensure correct MIME type for JavaScript modules
		if strings.Contains(resp.Request.URL.Path, ".js") {
			resp.Header.Set("Content-Type", "application/javascript")
		}
		return nil
	}

	return &simpleServer{
		addr:  addr,
		proxy: proxy,
	}
}

type LoadBalancer struct {
	port            string
	roundRobinCount int
	servers         []Server
	mu              sync.Mutex
}

func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:            port,
		roundRobinCount: 0,
		servers:         servers,
	}
}

func (s *simpleServer) Address() string {
	return s.addr
}

func (s *simpleServer) IsAlive() bool {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Get(s.addr)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK || resp.StatusCode < 400
}

func (s *simpleServer) Serve(rw http.ResponseWriter, req *http.Request) {
	s.proxy.ServeHTTP(rw, req)
}

func (lb *LoadBalancer) getNextAvailableServer() Server {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	for i := 0; i < len(lb.servers); i++ {
		server := lb.servers[lb.roundRobinCount%len(lb.servers)]
		lb.roundRobinCount++
		if server.IsAlive() {
			return server
		}
	}
	log.Fatal("No available servers")
	return nil
}

func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, req *http.Request) {
	targetServer := lb.getNextAvailableServer()
	fmt.Printf("Forwarding request to %s\n", targetServer.Address())
	targetServer.Serve(rw, req)
}

func main() {
	servers := []Server{
		// Add your server URLs here
		newSimpleServer("https://www.instagram.com"),
		newSimpleServer("https://www.facebook.com"),
		newSimpleServer("https://college-cart.vercel.app"),
		newSimpleServer("https://ayushranjan.tech"),
 // Example local server
	}

	lb := NewLoadBalancer("8000", servers)

	http.HandleFunc("/", lb.serveProxy)

	fmt.Printf("Serving on localhost:%s\n", lb.port)
	log.Fatal(http.ListenAndServe("localhost:"+lb.port, nil))
}