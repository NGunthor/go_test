package proxy

import (
	"math/rand"
	"time"
)

type httpServer struct {
	application Server
}

// HandleRequest handles
func (hs *httpServer) HandleRequest(url, method string) (int, string) {
	allowed := NewBoolGen(rand.NewSource(time.Now().UnixNano())).Bool()
	if !allowed {
		return 403, "Not Allowed"
	}
	return hs.application.HandleRequest(url, method)
}

// NewHttpServer ...
func NewHttpServer(s Server) Server {
	return &httpServer{s}
}
