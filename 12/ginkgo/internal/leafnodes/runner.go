package leafnodes

import (
	"time"

	"github.com/gy-kim/2020-golang-practice/12/ginkgo/internal/types"
)

type runner struct {
	isAsync          bool
	asyncFunc        func(chan<- interface{})
	syncFunc         func()
	codeLocation     types.CodeLocation
	timeoutThreshold time.Duration
	nodeType         types.SpecComponentType
	componentIndex   int
}
