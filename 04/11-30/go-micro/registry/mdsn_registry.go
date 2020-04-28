package registry

var (
	// use a .micro domain rather than .local
	mdnsDomain = "micro"
)

func newRegistry(opts ...Option) Registry {
	return nil
}

func NewRegistry(opts ...Option) Registry {
	return newRegistry(opts...)
}
