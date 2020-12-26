package suite

import (
	"github.com/gy-kim/2020-golang-practice/12/ginkgo/internal/containernode"
	"github.com/gy-kim/2020-golang-practice/12/ginkgo/internal/failer"
	"github.com/gy-kim/2020-golang-practice/12/ginkgo/internal/leafnodes"
	"github.com/gy-kim/2020-golang-practice/12/ginkgo/internal/specrunner"
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
	currentContainer  *containernode.ContainerNode

	deferredContainerNodes []deferredContainerNode

	containerIndex      int
	beforeSuiteNode     leafnodes.SuiteNode
	afterSuiteNode      leafnodes.SuiteNode
	runner              *specrunner.SpecRunner
	failer              *failer.Failer
	running             bool
	expandTopLevelNodes bool
}

func New(failer *failer.Failer) *Suite {
	topLevelContainer := containernode.New("[Top Level]", types.FlagTypeNone, types.CodeLocation{})

	return &Suite{
		topLevelContainer:      topLevelContainer,
		currentContainer:       topLevelContainer,
		failer:                 failer,
		containerIndex:         1,
		deferredContainerNodes: []deferredContainerNode{},
	}
}
