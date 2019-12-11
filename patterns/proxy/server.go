package proxy

type server interface {
	HandleRequest(string, string) (int, string)
}
