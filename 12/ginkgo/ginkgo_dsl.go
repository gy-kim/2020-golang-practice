package ginkgo

import (
	"flag"
	"io"
	"os"
	"time"

	"github.com/gy-kim/2020-golang-practice/12/ginkgo/config"
	"github.com/gy-kim/2020-golang-practice/12/ginkgo/internal/global"
	"github.com/gy-kim/2020-golang-practice/12/ginkgo/internal/writer"
)

const GINKGO_VERSION = config.VERSION
const GINKGO_PANIC = `
Your test failed.
Ginkgo panics to prevent subsequent assertions from running.
Normally Ginkgo rescues this panic so you shouldn't see it.

But, if you make an assertion in a goroutines, Ginkgo can't capture the panic.
To circumvent this, you should call

	defer GinkgoRecover()

at the top of the goroutine that caused this panic
`

func init() {
	config.Flags(flag.CommandLine, "ginkgo", true)
	GinkgoWriter = writer.New(os.Stdout)
}

//GinkgoWriter implements an io.Writer
//When running in verbose mode any writes to GinkgoWriter will be immediately printed
//to stdout. Otherwise, GinkgoWriter will buffer any writes produced during the current test and flush them to screen
// only if the current test fails.
var GinkgoWriter io.Writer

//The interface by which Ginkgo receives *testing.T
type GinkgoTestingT interface {
	Fail()
}

// GinkgoRandomSeed returns the seed to randomize spec execution order. It is
// useful for seeding your own pseudorandom number generator (PRNGs)
func GinkgoRandomSeed() int64 {
	return config.GinkgoConfig.RandomSeed
}

// GinkgoParallelNode returns the parallel node number for the current ginkgo process
// The node number is 1-indexed
func GinkgoParallelNode() int {
	return config.GinkgoConfig.ParallelNode
}

// Some matcher libraries or legacy codebases require a *testing.T
// func GinkgoT(optionalOffset ...int) GinkgoTInterface {

// }

// The interface returned by GinkgoT(). This covers most of the methods
// in the testing package's T.
type GinkgoTInterface interface {
	Cleanup(func())
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fail()
	FailNow()
	Failed() bool
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Helper()
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	Name() string
	Skip(args ...interface{})
	SkipNow()
	Skipf(format string, args ...interface{})
	Skipped() bool
	TempDir() string
}

///
//
//

func parseTimeout(timeout ...float64) time.Duration {
	if len(timeout) == 0 {
		return global.DefaultTimeout
	}
	return time.Duration(timeout[0] * float64(time.Second))
}
