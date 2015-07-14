package binding

import "net/http"

type Binding map[string]interface{}

func GetDefault(r *http.Request) Binding {
	return Binding{
		"PageTitle":  "Portafolio Nacho",
		"CurrentURL": r.URL.Path,
	}
}
