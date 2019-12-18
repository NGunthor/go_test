package chain_of_resp

//An interface for the chain
type Handler interface {
	Request(message string) string
}
