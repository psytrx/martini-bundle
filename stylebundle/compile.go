package stylebundle

import (
	"github.com/fjdumont/martini-bundle/concat"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	minifyService = "http://cssminifier.com/raw"
)

func compile(opts *Options, files ...string) string {
	css := concat.Files(files...)

	if opts.Minify {
		css = minify(css)
	}

	return css
}

func minify(css string) string {
	post := url.Values{
		"input": {css},
	}

	res, err := http.PostForm(minifyService, post)
	if err != nil {
		panic(err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		panic(err)
	}

	return string(resBody)
}
