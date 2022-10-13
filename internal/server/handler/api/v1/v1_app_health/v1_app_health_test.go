package v1_app_health_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vladimirok5959/golang-ip2location/internal/server/handler/api/v1/v1_app_health"
	"github.com/vladimirok5959/golang-ip2location/internal/server/handler/base"
)

var _ = Describe("Server", func() {
	Context("Endpoint", func() {
		var srv *httptest.Server
		var client *http.Client

		AfterEach(func() {
			srv.Close()
		})

		Context("/api/v1/app/health", func() {
			apiEndpoint := "/api/v1/app/health"

			BeforeEach(func() {
				handler := v1_app_health.Handler{Handler: base.Handler{Client: nil, Shutdown: nil}}
				srv = httptest.NewServer(handler)
				client = srv.Client()
			})

			AfterEach(func() {
				srv.Close()
			})

			It("respond with correct json", func() {
				resp, err := client.Get(srv.URL + apiEndpoint)
				Expect(err).To(Succeed())
				defer resp.Body.Close()

				Expect(resp.StatusCode).To(Equal(http.StatusOK))
				Expect(resp.Header.Get("Content-Type")).To(Equal("application/json"))

				body, err := io.ReadAll(resp.Body)
				Expect(err).To(Succeed())

				Expect(string(body)).To(MatchJSON(`{"health":"green"}`))
			})
		})
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Server")
}
