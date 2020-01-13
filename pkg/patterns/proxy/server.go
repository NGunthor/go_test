package proxy

// Server provides interface for server
type Server interface {
	HandleRequest(url, method string) (int, string)
}
