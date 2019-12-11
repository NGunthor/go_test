package proxy

type httpServer struct {
	application *application
}

func NewHttpServer() *httpServer {
	return &httpServer{&application{}}
}

func (hs *httpServer) HandleRequest(url, method string) (int, string) {
	allowed := NewBoolGen().Bool()
	if !allowed {
		return 403, "Not Allowed"
	}
	return hs.application.HandleRequest(url, method)
}


