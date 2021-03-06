package github_test

import (
	"testing"

	"go.uber.org/goleak"

	. "github.com/bartoszmajsak/github-changelog-generator/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestVersion(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecWithJUnitReporter(t, "Version Suite")
}

var _ = SynchronizedAfterSuite(func() {}, func() {
	goleak.VerifyNone(GinkgoT(),
		goleak.IgnoreTopFunction("github.com/bartoszmajsak/github-changelog-generator/vendor/github.com/onsi/ginkgo/internal/specrunner.(*SpecRunner).registerForInterrupts"),
	)
})
