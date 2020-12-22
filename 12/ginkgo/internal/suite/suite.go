package suite

import (
	"github.com/gy-kim/2020-golang-practice/12/ginkgo/internal/containernode"
	"github.com/gy-kim/2020-golang-practice/12/ginkgo/internal/leafnodes"
	"github.com/gy-kim/2020-golang-practice/12/ginkgo/internal/types"
)

type ginkgoTestingT interface {
	Fail()
}

type deferredContainerNode struct {
	text         string
	body         func()
	flag         types.FlagType
	codeLocation types.CodeLocation
}

type Suite struct {
	topLevelContainer *containernode.ContainerNode
	currrentContainer *containernode.ContainerNode

	deferredContainerNode []deferredContainerNode

	containerIndex  int
	beforeSuiteNode leafnodes.SuiteNode
}
