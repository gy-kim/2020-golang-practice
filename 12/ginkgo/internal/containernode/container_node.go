package containernode

import (
	"github.com/gy-kim/2020-golang-practice/12/ginkgo/internal/leafnodes"
	"github.com/gy-kim/2020-golang-practice/12/ginkgo/internal/types"
)

type subjectOrContainerNode struct {
	containerNode *ContainerNode
	subjectNode   leafnodes.SubjectNode
}

func (n subjectOrContainerNode) text() string {
	if n.containerNode != nil {
		return n.containerNode.Text()
	} else {
		return n.subjectNode.Text()
	}
}

type CollatedNodes struct {
	Containers []*ContainerNode
	Subject    leafnodes.SubjectNode
}

type ContainerNode struct {
	text         string
	flag         types.FlagType
	codeLocation types.CodeLocation

	setupNodes               []leafnodes.BasicNode
	subjectAndContainerNodes []subjectOrContainerNode
}

func New(text string, flag types.FlagType, codeLocation types.CodeLocation) *ContainerNode {
	return &ContainerNode{
		text:         text,
		flag:         flag,
		codeLocation: codeLocation,
	}
}

func (node *ContainerNode) Text() string {
	return node.text
}
