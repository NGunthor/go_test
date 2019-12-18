package proxy

import (
	"math/rand"
	"tests/patterns/proxy/application"
	"time"
)

type httpServer struct {
	application Server
}

func (hs *httpServer) HandleRequest(url, method string) (int, string) {
	allowed := NewBoolGen(rand.NewSource(time.Now().UnixNano())).Bool()
	if !allowed {
		return 403, "Not Allowed"
	}
	return hs.application.HandleRequest(url, method)
}

func main() {
	NewHttpServer(application.NewApplication())
}

func NewHttpServer(s Server) Server {
	return &httpServer{s}
}
