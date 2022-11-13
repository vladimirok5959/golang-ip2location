package base

import (
	"context"
	"html/template"
	"net/http"

	"github.com/vladimirok5959/golang-ip2location/internal/client"
)

type Handler struct {
	Client   *client.Client
	Shutdown context.CancelFunc
}

type ServerData struct {
	Additional    interface{}
	AppVersion    string
	AssetsVersion int
	URL           string
	WebURL        string
}

func (h Handler) FuncMap(w http.ResponseWriter, r *http.Request) template.FuncMap {
	return template.FuncMap{}
}
