package leafnodes

import "github.com/gy-kim/2020-golang-practice/12/ginkgo/internal/types"

type BasicNode interface {
	Type() types.SpecComponentType
	Run() (types.SpecState, types.SpecFailure)
	CodeLocation() types.CodeLocation
}

type SubjectNode interface {
	BasicNode

	Text() string
	Flag() types.FlagType
	Samples() int
}
