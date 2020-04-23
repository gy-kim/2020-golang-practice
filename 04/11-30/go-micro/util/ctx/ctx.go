package ctx

import (
	"context"
	"net/http"
	"strings"

	"github.com/gy-kim/2020-golang-practice/04/11-30/go-micro/metadata"
)

func FromRequest(r *http.Request) context.Context {
	ctx := context.Background()
	md := make(metadata.Metadata)
	for k, v := range r.Header {
		md[k] = strings.Join(v, ",")
	}
	return metadata.NewContext(ctx, md)
}
