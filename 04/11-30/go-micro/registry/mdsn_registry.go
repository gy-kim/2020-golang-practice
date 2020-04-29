package registry

var (
	// use a .micro domain rather than .local
	mdnsDomain = "micro"
)

type mdnsTxt struct {
	Service   string
	Version   string
	Endpoints []*Endpoint
	Metadata  map[string]string
}

func newRegistry(opts ...Option) Registry {
	return nil
}

func NewRegistry(opts ...Option) Registry {
	return newRegistry(opts...)
}
