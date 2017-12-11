package service_test

import (
	"perScoreServer/app/service"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Decrypt", func() {
	key := []byte("fkzfgk0FY2CaYJhyXbshnPJaRrFtCwfj")
	token := "VJmquZhZIqdhwnVwIyS6MtFrhse4DRSzprf4odxquiUFrOLfRA=="
	actual := service.Decrypt(key, token)
	expected := "encrypt this going on"
	It("should generate token", func() {
		Expect(actual).To(Equal(expected))
	})
})
