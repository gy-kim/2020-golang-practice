package suite

import "github.com/gy-kim/2020-golang-practice/12/ginkgo/internal/types"

type ginkgoTestingT interface {
	Fail()
}

type deferredContainerNode struct {
	text         string
	body         func()
	flag         types.FlagType
	codeLocation types.CodeLocation
}
