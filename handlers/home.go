package handlers

import (
	"net/http"

	"github.com/ignaciocarvajal/go-portafolio/models/binding"

	"github.com/unrolled/render"
	"github.com/zenazn/goji/web"
)

func GetHome(c web.C, w http.ResponseWriter, r *http.Request) {
	template := c.Env["render"].(*render.Render)
	bnd := binding.GetDefault(r)

	template.HTML(w, http.StatusOK, "index", bnd)
}
