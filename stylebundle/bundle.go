package stylebundle

import (
	"github.com/codegangsta/martini"
	"net/http"
)

type Options struct {
	Minify bool
}

func Default(files ...string) martini.Handler {
	opts := &Options{
		Minify: true,
	}

	return Bundle(opts, files...)
}

func Bundle(opts *Options, files ...string) martini.Handler {
	css := compile(opts, files...)

	return func(res http.ResponseWriter) string {
		res.Header().Set("Content-Type", "text/css")
		return css
	}
}
