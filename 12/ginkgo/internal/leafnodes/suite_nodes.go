package leafnodes

import "github.com/gy-kim/2020-golang-practice/12/ginkgo/internal/types"

type SuiteNode interface {
	Run(parallelNode int, parallelTotal int, syncHost string) bool
	Passed() bool
	Summary() *types.SetupSummary
}

type simpleSuiteNode struct {
}
