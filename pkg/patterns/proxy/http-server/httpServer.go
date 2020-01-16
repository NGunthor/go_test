package http_server

type Server interface {
	HandleRequest(string, string) (int, string)
}

type httpServer struct {
}

func (a *httpServer) HandleRequest(url, method string) (int, string) {
	if url == "/myProfile" && method == "GET" {
		return 200, "OK"
	}

	//...

	if url == "/create/product" && method == "POST" {
		return 201, "Product Created"
	}

	return 404, "Not Found"
}

func NewHttpServer() Server{
	return &httpServer{}
}
