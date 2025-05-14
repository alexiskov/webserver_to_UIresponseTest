package webserver

import (
	"fmt"
	"net/http"
	"os"
)

type WebServer struct {
	Driver *http.Server
}

// Start web server
func Init(port int) {
	srv := WebServer{
		Driver: &http.Server{
			Addr: fmt.Sprintf(":%d", port),
		},
	}
	http.HandleFunc("/", router)
	srv.Driver.ListenAndServe()
}

func router(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("ssid")
	if err != nil {
		return
	}

	switch r.URL.String() {
	case "/":
		if d, err := os.ReadFile("./main.html"); err != nil {
			http.Error(w, "Page is not found", http.StatusNotFound)
		} else {
			http.SetCookie(w, &http.Cookie{
				Name:     "ssid",
				Value:    "testCookie_01110",
				HttpOnly: true,
			})
			w.Write(d)
		}
	case "/test_js_request":
		d, err := os.ReadFile("./jsScripts/main.js")
		if err != nil {
			http.Error(w, "Page is not found <main>", http.StatusNotFound)
			return
		}
		w.Write(d)
	default:
		http.Error(w, "Page is not found <main>", http.StatusNotFound)
	}
}
