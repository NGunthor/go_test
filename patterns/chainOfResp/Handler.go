package chainOfResp

//An interface for the chain
type Handler interface {
	Request(message string) string
}
