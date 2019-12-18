package proxy

type application struct {
}

func (a *application) HandleRequest(url, method string) (int, string) {
	if url == "/myProfile" && method == "GET" {
		return 200, "OK"
	}

	//...

	if url == "/create/product" && method == "POST" {
		return 201, "Product Created"
	}

	return 404, "Not Found"
}

func NewApplication() *application{
	return &application{}
}
