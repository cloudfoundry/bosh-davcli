package cmd_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cloudfoundry/bosh-davcli/cmd"

	testcmd "github.com/cloudfoundry/bosh-davcli/cmd/testing"
	davconf "github.com/cloudfoundry/bosh-davcli/config"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
)

func runDelete(config davconf.Config, args []string) error {
	logger := boshlog.NewLogger(boshlog.LevelNone)
	factory := NewFactory(logger)
	factory.SetConfig(config) //nolint:errcheck

	cmd, err := factory.Create("delete")
	Expect(err).ToNot(HaveOccurred())

	return cmd.Run(args)
}

var _ = Describe("DeleteCmd", func() {
	var (
		handler       func(http.ResponseWriter, *http.Request)
		requestedBlob string
		ts            *httptest.Server
		config        davconf.Config
	)

	BeforeEach(func() {
		requestedBlob = "0ca907f2-dde8-4413-a304-9076c9d0978b"

		handler = func(w http.ResponseWriter, r *http.Request) {
			req := testcmd.NewHTTPRequest(r)

			username, password, err := req.ExtractBasicAuth()
			Expect(err).ToNot(HaveOccurred())
			Expect(req.URL.Path).To(Equal("/0d/" + requestedBlob))
			Expect(req.Method).To(Equal("DELETE"))
			Expect(username).To(Equal("some user"))
			Expect(password).To(Equal("some pwd"))

			w.WriteHeader(http.StatusOK)
		}
	})

	AfterEach(func() {
		ts.Close()
	})

	AssertDeleteBehavior := func() {
		It("with valid args", func() {
			err := runDelete(config, []string{requestedBlob})
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns err with incorrect arg count", func() {
			err := runDelete(davconf.Config{}, []string{})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Incorrect usage"))
		})
	}

	Context("with http endpoint", func() {
		BeforeEach(func() {
			ts = httptest.NewServer(http.HandlerFunc(handler))
			config = davconf.Config{
				User:     "some user",
				Password: "some pwd",
				Endpoint: ts.URL,
			}

		})

		AssertDeleteBehavior()
	})

	Context("with https endpoint", func() {
		BeforeEach(func() {
			ts = httptest.NewTLSServer(http.HandlerFunc(handler))

			rootCa, err := testcmd.ExtractRootCa(ts)
			Expect(err).ToNot(HaveOccurred())

			config = davconf.Config{
				User:     "some user",
				Password: "some pwd",
				Endpoint: ts.URL,
				TLS: davconf.TLS{
					Cert: davconf.Cert{
						CA: rootCa,
					},
				},
			}
		})

		AssertDeleteBehavior()
	})
})
