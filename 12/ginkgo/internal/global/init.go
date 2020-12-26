package global

import (
	"time"

	"github.com/gy-kim/2020-golang-practice/12/ginkgo/internal/failer"
	"github.com/gy-kim/2020-golang-practice/12/ginkgo/internal/suite"
)

const DefaultTimeout = time.Duration(1 * time.Second)

var Suite *suite.Suite
var Failer *failer.Failer

func init() {
	InitalizeGlobals()
}

func InitalizeGlobals() {
	Failer = failer.New()
	Suite = suite.New(Failer)
}
