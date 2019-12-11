package proxy

type Doer interface {
	Do(s string) string
}

type proxy struct {
	realObject Doer
}

type object struct {
	value string
}

func (p *proxy)Do(s string) string {
	if p.realObject == nil {
		p.realObject = &object{}
	}
	return "proxy " + p.realObject.Do(s)
}

func NewProxy() *proxy{
	return &proxy{}
}

func (o *object)Do(s string) string {
	return "value " + s
}


