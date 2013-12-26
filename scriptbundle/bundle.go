package scriptbundle

import (
	"github.com/codegangsta/martini"
	"net/http"
)

type Options struct {
	Wrap   bool
	Minify bool
}

func Default(files ...string) martini.Handler {
	opts := &Options{
		Wrap:   true,
		Minify: true,
	}

	return Bundle(opts, files...)
}

func Bundle(opts *Options, files ...string) martini.Handler {
	content := compile(opts, files...)

	return func(res http.ResponseWriter) string {
		res.Header().Set("Content-Type", "text/javascript")
		return content
	}
}
