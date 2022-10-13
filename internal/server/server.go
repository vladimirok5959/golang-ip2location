package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/vladimirok5959/golang-ip2location/internal/client"
	"github.com/vladimirok5959/golang-ip2location/internal/consts"
	"github.com/vladimirok5959/golang-ip2location/internal/server/handler/api/v1/v1_app_health"
	"github.com/vladimirok5959/golang-ip2location/internal/server/handler/api/v1/v1_ip2location"
	"github.com/vladimirok5959/golang-ip2location/internal/server/handler/base"
	"github.com/vladimirok5959/golang-ip2location/internal/server/handler/page/page_index"
	"github.com/vladimirok5959/golang-utils/utils/http/apiserv"
	"github.com/vladimirok5959/golang-utils/utils/http/helpers"
)

func NewMux(ctx context.Context, shutdown context.CancelFunc, client *client.Client) *apiserv.ServeMux {
	mux := apiserv.NewServeMux()

	handler := base.Handler{
		Client:   client,
		Shutdown: shutdown,
	}

	// Pages
	mux.Get("/", page_index.Handler{Handler: handler})

	// API
	mux.Get("/api/v1/app/health", v1_app_health.Handler{Handler: handler})
	mux.Get("/api/v1/app/status", helpers.HandleAppStatus())
	mux.Get("/api/v1/ip2location/{s}", v1_ip2location.Handler{Handler: handler})

	return mux
}

func New(ctx context.Context, shutdown context.CancelFunc, client *client.Client) (*http.Server, error) {
	mux := NewMux(ctx, shutdown, client)
	srv := &http.Server{
		Addr:    consts.Config.Host + ":" + consts.Config.Port,
		Handler: mux,
	}
	go func() {
		fmt.Printf("Web server: http://%s:%s/\n", consts.Config.Host, consts.Config.Port)
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				fmt.Printf("Web server startup error: %s\n", err.Error())
				shutdown()
				return
			}
		}
	}()
	return srv, nil
}
