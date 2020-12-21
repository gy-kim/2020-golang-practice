package ginkgo

import (
	"github.com/gy-kim/2020-golang-practice/12/ginkgo/config"
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

}
