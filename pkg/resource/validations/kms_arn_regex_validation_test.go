package validations

import (

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Validations", func() {
	Describe("validateKMSKeyARN", func() {
		var (
			kmsKeyARN  string
		)

		BeforeEach(func() {
			kmsKeyARN = ""
		})

		Context("when kmsKeyARN is not empty and matches the regex", func() {
			BeforeEach(func() {
				kmsKeyARN = "arn:aws:kms:us-east-1:111111111111:key/mrk-0123456789abcdef0123456789abcdef"
			})

			It("should not return an error", func() {
				err := ValidateKMSKeyARN(&kmsKeyARN)
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("when kmsKeyARN is not empty and does not match the regex", func() {
			BeforeEach(func() {
				kmsKeyARN = "invalid-kms-key-arn"
			})

			It("should return an error", func() {
				err := ValidateKMSKeyARN(&kmsKeyARN)
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
