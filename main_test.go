package main_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Main", func() {
	Describe("when run with --help", func() {
		It("return exit code 0 with help message", func() {
			command := exec.Command(pathCLI, "--help")
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ShouldNot(HaveOccurred())
			Eventually(session).Should(gexec.Exit(0))
			Eventually(session.Out).Should(gbytes.Say("Usage:"))
			Eventually(session.Err.Contents()).Should(BeEmpty())
		})
	})
})