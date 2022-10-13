package v1_ip2location

import (
	"net/http"

	"github.com/vladimirok5959/golang-ip2location/internal/client"
	"github.com/vladimirok5959/golang-ip2location/internal/server/handler/base"
	"github.com/vladimirok5959/golang-utils/utils/http/apiserv"
	"github.com/vladimirok5959/golang-utils/utils/http/helpers"
	"github.com/vladimirok5959/golang-utils/utils/http/render"
)

type Handler struct {
	base.Handler
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var res *client.Result
	var err error

	if h.Client != nil {
		if res, err = h.Client.IP2Location(r.Context(), apiserv.GetParams(r)[1].String()); err != nil {
			helpers.RespondAsBadRequest(w, r, err)
			return
		}
	}

	if !render.JSON(w, r, res) {
		return
	}
}
