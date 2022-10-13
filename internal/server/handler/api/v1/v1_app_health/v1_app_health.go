package v1_app_health

import (
	"net/http"

	"github.com/vladimirok5959/golang-ip2location/internal/server/handler/base"
	"github.com/vladimirok5959/golang-utils/utils/http/render"
)

type HealthColor int64

const (
	HealthColorGreen HealthColor = iota
	HealthColorOrange
	HealthColorRed
)

func (c HealthColor) String() string {
	switch c {
	case HealthColorGreen:
		return "green"
	case HealthColorOrange:
		return "orange"
	case HealthColorRed:
		return "red"
	}
	return "unknown"
}

type Handler struct {
	base.Handler
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	colorHealth := HealthColorGreen

	var resp struct {
		Health string `json:"health"`
	}

	resp.Health = colorHealth.String()

	if !render.JSON(w, r, resp) {
		return
	}
}
