package page_index

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/vladimirok5959/golang-ip2location/internal/consts"
	"github.com/vladimirok5959/golang-ip2location/internal/server/handler/base"
	"github.com/vladimirok5959/golang-ip2location/internal/server/web"
	"github.com/vladimirok5959/golang-utils/utils/http/helpers"
	"github.com/vladimirok5959/golang-utils/utils/http/render"
)

type Handler struct {
	base.Handler
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := &base.ServerData{
		URL:    r.URL.Path,
		WebURL: consts.Config.WebURL,
	}

	var additional struct {
		ClientIP  string
		GeoIPData template.HTML
	}

	additional.ClientIP = helpers.ClientIP(r)

	if h.Client != nil {
		if res, err := h.Client.IP2Location(r.Context(), additional.ClientIP); err == nil {
			if j, err := json.Marshal(res); err == nil {
				additional.GeoIPData = template.HTML(string(j))
			}
		}
	}

	data.Additional = additional

	if !render.HTML(w, r, h.FuncMap(w, r), data, web.IndexHtml, http.StatusOK) {
		return
	}
}
