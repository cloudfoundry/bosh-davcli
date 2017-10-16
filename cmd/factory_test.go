package cmd_test

import (
	"crypto/x509"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cloudfoundry/bosh-davcli/cmd"
	testcmd "github.com/cloudfoundry/bosh-davcli/cmd/testing"
	davconf "github.com/cloudfoundry/bosh-davcli/config"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
)

func buildFactory() (factory Factory) {
	config := davconf.Config{User: "some user"}
	logger := boshlog.NewLogger(boshlog.LevelNone)
	factory = NewFactory(logger)
	factory.SetConfig(config)
	return
}

func init() {
	Describe("Testing with Ginkgo", func() {
		It("factory create a put command", func() {
			factory := buildFactory()
			cmd, err := factory.Create("put")

			Expect(err).ToNot(HaveOccurred())
			Expect(reflect.TypeOf(cmd)).To(Equal(reflect.TypeOf(PutCmd{})))
		})

		It("factory create a get command", func() {
			factory := buildFactory()
			cmd, err := factory.Create("get")

			Expect(err).ToNot(HaveOccurred())
			Expect(reflect.TypeOf(cmd)).To(Equal(reflect.TypeOf(GetCmd{})))
		})

		It("factory create a delete command", func() {
			factory := buildFactory()
			cmd, err := factory.Create("delete")

			Expect(err).ToNot(HaveOccurred())
			Expect(reflect.TypeOf(cmd)).To(Equal(reflect.TypeOf(DeleteCmd{})))
		})

		It("factory create when cmd is unknown", func() {
			factory := buildFactory()
			_, err := factory.Create("some unknown cmd")

			Expect(err).To(HaveOccurred())
		})

		It("get command uses provided test ca cert and errors out due to mismatch", func() {
			factory := buildFactory()

			requestedBlob := "0ca907f2-dde8-4413-a304-9076c9d0978b"
			targetFilePath := filepath.Join(os.TempDir(), "testRunGetCommand.txt")
			handler := func(w http.ResponseWriter, r *http.Request) {
				req := testcmd.NewHTTPRequest(r)

				username, password, err := req.ExtractBasicAuth()
				Expect(err).ToNot(HaveOccurred())
				Expect(req.URL.Path).To(Equal("/0d/" + requestedBlob))
				Expect(req.Method).To(Equal("GET"))
				Expect(username).To(Equal("some user"))
				Expect(password).To(Equal("some pwd"))

				w.Write([]byte("this is your blob"))
			}

			ts := httptest.NewTLSServer(http.HandlerFunc(handler))
			defer ts.Close()

			factoryConfig := davconf.Config{
				User:          "some user",
				Password:      "some pwd",
				Endpoint:      ts.URL,
				RetryAttempts: 3,
				CA:            "ca cert",
			}

			expectedCaCertPool := x509.NewCertPool()
			expectedCaCertPool.AppendCertsFromPEM([]byte(factoryConfig.CA))
			factory.SetConfig(factoryConfig)

			cmd, err := factory.Create("get")
			Expect(err).ToNot(HaveOccurred())

			err = cmd.Run([]string{requestedBlob, targetFilePath})
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("x509: certificate signed by unknown authority")))
		})
	})
}
