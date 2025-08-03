package routes

import "net/http"

type HttpServer struct{}

func (h *HttpServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "text/html")
	documento := "Kevin Morales"
	rw.Write([]byte(documento))
}
