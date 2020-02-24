package main

import (
	"net/http"
	"log"
	"github.com/julienschmidt/httprouter"
	"strings"
)

func main() {
	startServer()
}

func startServer() {
	activeRouter := httprouter.New()
	activeRouter.POST("/previewCard", previewCard)

	static := httprouter.New()
	static.ServeFiles("/frontEnd/*filepath", neuteredFileSystem{http.Dir("./frontEnd")})

	activeRouter.NotFound = static

	server := &http.Server{
		Addr:    ":12345",
		Handler: RequestLogger{activeRouter},
	}

	log.Fatal(server.ListenAndServe())
}

type RequestLogger struct {
	h http.Handler
	//l *Logger -- somehow connect this to the log logger type
}

func (rl RequestLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//start := 123
	log.Printf("Started %s %s", r.Method, r.URL.Path)
	rl.h.ServeHTTP(w, r)
	log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, 123)
}

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {

	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := strings.TrimSuffix(path, "/") + "/index.html"
		if _, err := nfs.fs.Open(index); err != nil {
			return nil, err
		}
	}

	return f, nil
}
