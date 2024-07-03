package ports

import (
	"net/http"
)

type HttpServer struct {}

func NewHttpServer() HttpServer {
	return HttpServer{}
}

func (h HttpServer) GetTrainings(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get trainings"))
}
