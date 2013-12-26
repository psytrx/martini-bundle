package scriptbundle

import (
	"fmt"
	"github.com/fjdumont/martini-bundle/concat"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	minifyService = "http://closure-compiler.appspot.com/compile"
)

func compile(opts *Options, files ...string) string {
	content := concat.Files(files...)

	if opts.Wrap {
		content = wrap(content)
	}
	if opts.Minify {
		content = minify(content)
	}

	return content
}

func wrap(script string) string {
	return fmt.Sprintf("(function () { %s })();", script)
}

func minify(script string) string {
	post := url.Values{
		"js_code":           {script},
		"compilation_level": {"SIMPLE_OPTIMIZATIONS"},
		"output_format":     {"text"},
		"output_info":       {"compiled_code"},
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
