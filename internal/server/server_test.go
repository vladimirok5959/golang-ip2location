package server_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vladimirok5959/golang-ip2location/internal/server"
)

var _ = Describe("Server", func() {
	Context("Endpoint", func() {
		var ctx = context.Background()
		var srv *httptest.Server
		var client *http.Client

		AfterEach(func() {
			srv.Close()
		})

		Context("Routes", func() {
			BeforeEach(func() {
				mux := server.NewMux(ctx, nil, nil)
				srv = httptest.NewServer(mux)
				client = srv.Client()
			})

			AfterEach(func() {
				srv.Close()
			})

			Context("Route", func() {
				It("must exists", func() {
					var routes = []string{
						// Pages
						"/",

						// API
						"/api/v1/app/health",
						"/api/v1/app/status",
						"/api/v1/ip2location/127.0.0.1",

						// Assets
						"/styles.css",
					}

					for _, route := range routes {
						resp, err := client.Get(srv.URL + route)
						resp.Body.Close()
						Expect(err).To(Succeed())
						Expect(resp.StatusCode).NotTo(Equal(http.StatusNotFound))
					}
				})

				It("must response with 404", func() {
					resp, err := client.Get(srv.URL + "/qwertyuiopasdfghjklzxcvbnm")
					resp.Body.Close()
					Expect(err).To(Succeed())
					Expect(resp.StatusCode).To(Equal(http.StatusNotFound))
				})
			})
		})
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Server")
}
