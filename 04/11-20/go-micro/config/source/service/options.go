package service

import (
	"context"

	"github.com/gy-kim/2020-golang-practice/04/11-20/go-micro/config/source"
)

type serviceNameKey struct{}
type namespaceKey struct{}
type pathKey struct{}

func ServiceName(name string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
	}
}

func Namespace(namespace string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, namespaceKey{}, namespace)
	}
}

func Path(path string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, pathKey{}, path)
	}
}
