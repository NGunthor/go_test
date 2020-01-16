package proxy

type boolGenerator interface {
	Generate() bool
}

type server interface {
	HandleRequest(string, string) (int, string)
}

type validator struct {
	application server
	generator boolGenerator
}

func (hs *validator) HandleRequest(url, method string) (int, string) {
	allowed := hs.generator.Generate()
	if !allowed {
		return 403, "Not Allowed"
	}
	return hs.application.HandleRequest(url, method)
}

func NewValidator(server server, generator boolGenerator) server {
	return &validator{
		application: server,
		generator:   generator,
	}
}
