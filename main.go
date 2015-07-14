package main

import (
	"github.com/ignaciocarvajal/go-portafolio/handlers"
	"github.com/ignaciocarvajal/go-portafolio/middleware/render"
	"github.com/theosomefactory/goji-static"
	"github.com/zenazn/goji"
)

func main() {

	goji.Use(static.Static("static"))
	goji.Use(render.InjectRender)

	goji.Get("/", handlers.GetHome)
	goji.Serve()
}
