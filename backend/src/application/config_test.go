package application_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/jarmex/ussd-designer/src/application"
)

var _ = Describe("Config", func() {
	BeforeEach(func() {
		os.Setenv("PORT", "8081")
		os.Setenv("REDIS_ADDRESS", "localhost:6379")
	})

	AfterEach(func() {
		os.Unsetenv("PORT")
	})

	It("should load server port from env", func() {
		Expect(application.LoadConfig().ServerPort).To(Equal(uint16(8081)))
	})

	It("should load redis server address", func() {
		Expect(application.LoadConfig().RedisAddress).To(Equal("localhost:6379"))
	})
})
