package page_index

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strings"

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
		AppVersion:    consts.AppVersion,
		AssetsVersion: consts.AssetsVersion,
		URL:           r.URL.Path,
		WebURL:        consts.Config.WebURL,
	}

	var additional struct {
		ClientIP  string
		GeoIPData template.HTML
	}

	additional.ClientIP = helpers.ClientIP(r)

	ip := strings.Trim(r.FormValue("ip"), " ")
	if ip != "" && len([]rune(ip)) <= 15 {
		additional.ClientIP = ip
	}

	if h.Client != nil {
		if res, err := h.Client.IP2Location(r.Context(), additional.ClientIP); err == nil {
			if j, err := json.MarshalIndent(res, "<br>", "&nbsp;&nbsp;"); err == nil {
				additional.GeoIPData = template.HTML(string(j))
			}
		}
	}

	data.Additional = additional

	if !render.HTML(w, r, h.FuncMap(w, r), data, web.IndexHtml, http.StatusOK) {
		return
	}
}
