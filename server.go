package main

import (
	"net/http"
	"regexp"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/lestrrat/go-xslate"
)

func main() {
	goji.Get("/hello/:name", hello)

	staticPattern := regexp.MustCompile("^/(css|js)")
	goji.Handle(staticPattern, http.FileServer(http.Dir("./static")))
	goji.Serve()
}

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	tx, _ := xslate.New(xslate.Args {
		"Loader": xslate.Args {
			"LoadPaths": []string { "./templates" },
		},
		"Parser": xslate.Args{"Syntax": "TTerse"},
	})

	tx.RenderInto(w, "hello.tt", xslate.Vars {
		"name": c.URLParams["name"],
	});
}

