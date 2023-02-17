package router

import (
	"fmt"
	"net/http"
)

// Router is the interface that wraps the basic ServeHTTP method.
type router struct {
	mux http.ServeMux
}

type httpMethod string

const (
	GET     httpMethod = "GET"
	POST    httpMethod = "POST"
	PUT     httpMethod = "PUT"
	DELETE  httpMethod = "DELETE"
	PATCH   httpMethod = "PATCH"
	HEAD    httpMethod = "HEAD"
	OPTIONS httpMethod = "OPTIONS"
)

func (r *router) addRoute(path string, method httpMethod, handle http.HandlerFunc) {
	internalHandler := func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != path {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if req.Method != string(method) {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		handle(w, req)
	}

	r.mux.HandleFunc(path, internalHandler)
}

// NewRouter returns a new Router instance.
func NewRouter() *router {
	return &router{
		mux: *http.NewServeMux(),
	}
}

func (r *router) ListenAndServe(host string, port int) error {
	addr := fmt.Sprintf("%s:%d", host, port)
	return http.ListenAndServe(addr, &r.mux)
}

func (r *router) GET(path string, handler http.HandlerFunc) {
	r.addRoute(path, GET, handler)
}

func (r *router) POST(path string, handler http.HandlerFunc) {
	r.addRoute(path, POST, handler)
}

func (r *router) PUT(path string, handler http.HandlerFunc) {
	r.addRoute(path, PUT, handler)
}

func (r *router) DELETE(path string, handler http.HandlerFunc) {
	r.addRoute(path, DELETE, handler)
}

func (r *router) PATCH(path string, handler http.HandlerFunc) {
	r.addRoute(path, PATCH, handler)
}

func (r *router) HEAD(path string, handler http.HandlerFunc) {
	r.addRoute(path, HEAD, handler)
}

func (r *router) OPTIONS(path string, handler http.HandlerFunc) {
	r.addRoute(path, OPTIONS, handler)
}
