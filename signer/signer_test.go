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
	duration := time.Duration(15 * time.Minute)
	timeStamp := time.Date(2019, 8, 26, 11, 11, 0, 0, time.UTC)

	Context("HMAC signature", func() {
		expected := "6140482f6d5846c158ff0f8dad88803d49c88657a716e2c46ba5ac41486918b2"

		It("generates a HMAC signature", func() {
			expirationTime := timeStamp.Add(duration)

			actual := signer.GenerateSignature(objectID, verb, timeStamp, expirationTime)
			Expect(actual).To(Equal(expected))
		})
	})

	Context("HMAC Signed URL", func() {
		path := "http://api.foo.bar/"
		expected := "http://api.foo.bar/signed/fake-object-id?st=6140482f6d5846c158ff0f8dad88803d49c88657a716e2c46ba5ac41486918b2&ts=1566817860&e=1566818760"

		It("Generates a properly formed URL", func() {
			actual, err := signer.GenerateSignedURL(path, objectID, verb, timeStamp, duration)

			Expect(err).To(BeNil())
			Expect(actual).To(Equal(expected))
		})
	})
})
