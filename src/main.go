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
	webMux := http.NewServeMux()

	webMux.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("src/static")),
		),
	)
	webMux.Handle(
		"/assets/",
		http.StripPrefix(
			"/assets/",
			http.FileServer(http.Dir("assets/")),
		),
	)

	webMux.HandleFunc("/", handlers.HomeHandler)
	webMux.HandleFunc("/design-system", handlers.DesignHandler)
	webMux.HandleFunc("/test", handlers.TestHandler)

	webMux.HandleFunc("/user/edit/", handlers.UserFormHandler)
	webMux.HandleFunc("/user/edit-cancel/", handlers.UserInfoHandler)

	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/api/user-data", handlers.UserDataHandler)
	apiMux.HandleFunc("/api/activity", handlers.ActivityHandler)
	apiMux.HandleFunc("/api/activity/", handlers.ActivityIdHandler)

	fmt.Println("Server running at http://localhost:" + config.ServerPort)

	// Try to find local IP
	addrs, err := net.InterfaceAddrs()
	if err == nil {
		for _, addr := range addrs {
			// check the address type and ignore loopback
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil { // IPv4
					fmt.Println(
						"Accessible on your LAN at: http://" + ipnet.IP.String() + ":" + config.ServerPort,
					)
				}
			}
		}
	}

	rootMux := http.NewServeMux()
	rootMux.Handle("/", webMux)
	rootMux.Handle("/api/", apiMux)

	http.ListenAndServe(":"+config.ServerPort, rootMux)
}
