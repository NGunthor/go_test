package chainOfResp

type Handler interface {
	Request(message string) string
	SetNext(handler Handler)
}
