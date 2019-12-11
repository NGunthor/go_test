package proxy

type Server interface {
	HandleRequest(string, string) (int, string)
}
