package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/fdanctl/p5r-stats/config"
	"github.com/fdanctl/p5r-stats/src/handlers"
	"github.com/fdanctl/p5r-stats/src/render"
)

func main() {
	render.Init()

	http.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("src/static")),
		),
	)
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/user-data", handlers.UserDataHandler)
	http.HandleFunc("/activity", handlers.ActivityHandler)
	http.HandleFunc("/activity/", handlers.ActivityIdHandler)
	http.HandleFunc("/design-system", handlers.DesignHandler)
	http.HandleFunc("/test", handlers.TestHandler)

	fmt.Println("Server running at http://localhost:" + config.ServerPort)

	// Try to find local IP
	addrs, err := net.InterfaceAddrs()
	if err == nil {
		for _, addr := range addrs {
			// check the address type and ignore loopback
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil { // IPv4
					fmt.Println("Accessible on your LAN at: http://" + ipnet.IP.String() + ":" + config.ServerPort)
				}
			}
		}
	}
	http.ListenAndServe(":"+config.ServerPort, nil)
}
