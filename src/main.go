package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/fdanctl/p5r-stats/config"
	"github.com/fdanctl/p5r-stats/src/handlers"
	"github.com/fdanctl/p5r-stats/src/middleware"
	"github.com/fdanctl/p5r-stats/src/render"
)

func main() {
	render.Init()

	webMux := http.NewServeMux() // returns full HTML page
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
	webMux.HandleFunc("/radar", handlers.RadarHandler)
	webMux.HandleFunc("/", handlers.HomeHandler)
	webMux.HandleFunc("/design-system", handlers.DesignHandler)
	webMux.HandleFunc("/test", handlers.TestHandler)

	partialsMux := http.NewServeMux() // returns HTMX fragment
	partialsMux.HandleFunc(
		"/partials/user-data",
		middleware.RequireHTMX(handlers.UserDataHandler),
	)
	partialsMux.HandleFunc(
		"/partials/user/edit/",
		middleware.RequireHTMX(handlers.UserFormHandler),
	)
	partialsMux.HandleFunc(
		"/partials/user/edit-cancel/",
		middleware.RequireHTMX(handlers.UserInfoHandler),
	)
	partialsMux.HandleFunc(
		"/partials/stat",
		middleware.RequireHTMX(handlers.StatHandler),
	)
	partialsMux.HandleFunc(
		"/partials/activity",
		middleware.RequireHTMX(handlers.ActivityHandler),
	)
	partialsMux.HandleFunc(
		"/partials/activity/",
		middleware.RequireHTMX(handlers.ActivityWithIdHandler),
	)
	partialsMux.HandleFunc(
		"/partials/settings",
		middleware.RequireHTMX(handlers.SettingsModalHandler),
	)

	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/api/health", handlers.HealthHandler)

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
	rootMux.Handle("/partials/", partialsMux)
	rootMux.Handle("/api/", apiMux)

	http.ListenAndServe(":"+config.ServerPort, rootMux)
}
