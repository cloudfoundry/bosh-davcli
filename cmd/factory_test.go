package cmd_test

import (
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cloudfoundry/bosh-davcli/cmd"
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
	})

}
