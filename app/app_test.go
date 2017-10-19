package app_test

import (
	"errors"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cloudfoundry/bosh-davcli/app"
	davconf "github.com/cloudfoundry/bosh-davcli/config"
)

type FakeRunner struct {
	Config  davconf.Config
	RunArgs []string
	RunErr  error
}

func (r *FakeRunner) SetConfig(newConfig davconf.Config) {
	r.Config = newConfig
}

func (r *FakeRunner) Run(cmdArgs []string) (err error) {
	r.RunArgs = cmdArgs
	return r.RunErr
}

func pathToFixture(file string) string {
	pwd, err := os.Getwd()
	Expect(err).ToNot(HaveOccurred())

	fixturePath := filepath.Join(pwd, "../test_assets", file)

	absPath, err := filepath.Abs(fixturePath)
	Expect(err).ToNot(HaveOccurred())

	return absPath
}

func init() {
	Describe("Testing with Ginkgo", func() {
		It("reads the CA cert from config", func() {
			runner := &FakeRunner{}

			app := New(runner)
			err := app.Run([]string{"dav-cli", "-c", pathToFixture("dav-cli-config-with-ca.json"), "put", "localFile", "remoteFile"})
			Expect(err).ToNot(HaveOccurred())

			expectedConfig := davconf.Config{
				User:     "some user",
				Password: "some pwd",
				Endpoint: "https://example.com/some/endpoint",
				CaCert:   "ca-cert",
			}

			Expect(runner.Config).To(Equal(expectedConfig))
			Expect(runner.Config.CaCert).ToNot(BeNil())
		})

		It("runs the put command", func() {
			runner := &FakeRunner{}

			app := New(runner)
			err := app.Run([]string{"dav-cli", "-c", pathToFixture("dav-cli-config.json"), "put", "localFile", "remoteFile"})
			Expect(err).ToNot(HaveOccurred())

			expectedConfig := davconf.Config{
				User:     "some user",
				Password: "some pwd",
				Endpoint: "http://example.com/some/endpoint",
			}

			Expect(runner.Config).To(Equal(expectedConfig))
			Expect(runner.Config.CaCert).To(BeEmpty())
			Expect(runner.RunArgs).To(Equal([]string{"put", "localFile", "remoteFile"}))
		})

		It("returns error with no config argument", func() {
			runner := &FakeRunner{}

			app := New(runner)
			err := app.Run([]string{"put", "localFile", "remoteFile"})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Config file arg `-c` is missing"))
		})
		It("prints the version info with the -v flag", func() {
			runner := &FakeRunner{}
			app := New(runner)
			err := app.Run([]string{"dav-cli", "-v"})
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns error from the cmd runner", func() {
			runner := &FakeRunner{
				RunErr: errors.New("fake-run-error"),
			}

			app := New(runner)
			err := app.Run([]string{"dav-cli", "-c", pathToFixture("dav-cli-config.json"), "put", "localFile", "remoteFile"})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-run-error"))
		})
	})
}
