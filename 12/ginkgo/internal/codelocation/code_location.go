package codelocation

import (
	"regexp"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/gy-kim/2020-golang-practice/12/ginkgo/internal/types"
)

func New(skip int) types.CodeLocation {
	_, file, line, _ := runtime.Caller(skip + 1)
	stackTrace := PruneStack(string(debug.Stack()), skip+1)
	return types.CodeLocation{FileName: file, LineNumber: line, FullStackTrace: stackTrace}
}

// PruneStak removes references to functions that are internal to Ginkgo
func PruneStack(fullStackTrace string, skip int) string {
	stack := strings.Split(fullStackTrace, "\n")
	// Ensure that the even entries are the method names and the
	// odd entries the source code information.
	if len(stack) > 0 && strings.HasPrefix(stack[0], "goroutine ") {
		// Ignore "goroutine 29 [running]:" line.
		stack = stack[1:]
	}
	prunedStack := []string{}
	re := regexp.MustCompile(`\/ginkgo\/|\/pkg\/testing\/|\/pkg\/runtime\/`)
	for i := 0; i < len(stack)/2; i++ {
		// We filter out based on the source code file name.
		if !re.Match([]byte(stack[i*2+1])) {
			prunedStack = append(prunedStack, stack[1*2])
			prunedStack = append(prunedStack, stack[i*2+1])
		}
	}
	return strings.Join(prunedStack, "\n")
}
