package signer_test

import (
	"github.com/cloudfoundry/bosh-davcli/signer"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Signer", func() {
	secret := "mefq0umpmwevpv034m890j34m0j0-9!fijm434j99j034mjrwjmv9m304mj90;2ef32buf32gbu2i3"
	objectID := "fake-object-id"
	verb := "get"
	signer := signer.NewSigner(secret)
	durationInSeconds := int64(600)
	startUnixTime := int64(1257894000)
	timeStamp := time.Unix(startUnixTime, 0)

	Context("HMAC signature", func() {
		expected := "c0558b9570aa91e528b9f527640bc42527e57eb26fb4b86ba0ba15e213504707"

		It("generates a HMAC signature", func() {
			expirationTime := time.Unix(startUnixTime+int64(durationInSeconds), 0)

			actual := signer.GenerateSignature(objectID, verb, timeStamp, expirationTime)
			Expect(actual).To(Equal(expected))
		})
	})

	Context("HMAC Signed URL", func() {
		path := "http://api.foo.bar/"
		expected := "http://api.foo.bar/fake-object-id?st=0994f5a3f1e602d9f5dbe87952f0cb27fc6f27e80a7ab879c37dfe343157ef8a&ts=1257894000&e=1257894600"

		It("Generates a properly formed URL", func() {
			actual, err := signer.GenerateSignedURL(path, objectID, verb, timeStamp, durationInSeconds)

			Expect(err).To(BeNil())
			Expect(actual).To(Equal(expected))
		})
	})
})
