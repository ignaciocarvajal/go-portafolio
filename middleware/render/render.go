package render

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/unrolled/render"
	"github.com/zenazn/goji/web"
)

func InjectRender(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		c.Env["render"] = render.New(render.Options{
			Directory:  "views",
			Layout:     "",
			Extensions: []string{".tmpl"},
			Funcs: []template.FuncMap{template.FuncMap{
				"classIfHere": func(path, class string) template.HTMLAttr {
					if r.URL.Path == path {
						return template.HTMLAttr(fmt.Sprintf(`class="%s"`, class))
					}
					return template.HTMLAttr("")
				},
				"date": func(date time.Time) string {
					return date.Format("January 2, 2006")
				},
				"checkIfInArray": func(x string, slice []string) template.HTMLAttr {
					for _, y := range slice {
						if x == y {
							return template.HTMLAttr("checked")
						}
					}
					return template.HTMLAttr("")
				},
			}},
		})

		h.ServeHTTP(w, r)

	}

	return http.HandlerFunc(fn)
}
