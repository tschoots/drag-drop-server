package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"golang.org/x/net/websocket"
)

// template struct can be used to load template pages
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}



func main() {
	var port = flag.String("port", ":8080", "the port the server will be listening on")
	flag.Parse()
	
	u := uploader{dir: "downloads"}
	
	http.Handle("/", &templateHandler{filename: "drag-drop.html"})
	http.Handle("/upload", websocket.Handler(u.UploadHandler))
	
	// Start the web server
	log.Println("Starting web server on", *port)
	if err := http.ListenAndServe(*port, nil); err != nil {
		log.Fatal("listenAndServe:", err)
	}
}
