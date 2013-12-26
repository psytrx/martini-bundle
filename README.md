# martini-bundle

`martini-bundle` adds asset optimization to **[Martini](https://github.com/codegangsta/martini)**.

## Build State

[![wercker status](https://app.wercker.com/status/1acc2e1acb54bd7cbef8ecf0c4193c1e "wercker status")](https://app.wercker.com/project/bykey/1acc2e1acb54bd7cbef8ecf0c4193c1e)

## Getting Started

Install all the packages (go 1.1 and greater is required):

    go get github.com/fjdumont/martini-bundle/...
  
and

~~~ go
import "github.com/fjdumont/martini-bundle/scriptbundle"
import "github.com/fjdumont/martini-bundle/stylebundle"
~~~

## Components

### scriptbundle

`martini-bundle/scriptbundle` adds **concatenation**, **scope wrapping** and **minification** to your JavaScript files.

- concatenation: simply concatenates all file contents
- scope wrapping: wraps an IIFE/IFFY/SIAF around the concatenated scripts
- minification: uses the [Closure Compiler Service](http://closure-compiler.appspot.com/home) REST API to minify your scripts

Simply use `scriptbundle.Default(files ...string)` as a `martini.Handler` to register your script files as a bundle:

~~~ go
m := martini.Classic()
m.Get("/js/app.js", scriptbundle.Default(
  "public/js/jquery.js",
  "public/js/app-utils.js",
  "public/js/app.js",
))
~~~

Alternatively, plug your own `scriptbundle.Options` into `scriptbundle.Bundle(opts *scriptbundle.Options, files ...string)` to customize your optimizations:

~~~ go
type Options struct {
  Wrap   bool
  Minify bool
}
~~~

### stylebundle

`martini-bundle/stylebundle` adds **concatenation** and **minification** to your Styleseets.

- concatenation: simply concatenates all file contents
- minification: uses the [CSS Minifier](http://cssminifier.com/) REST API to minify your styles

Use `stylebundle.Default(files ...string)` as a `martini.Handler` to compile your styles into a bundle and use them:

~~~ go
m := martini.Classic()
m.Get("/css/app.css", stylebundle.Default(
  "public/js/vendor/bootstrap.css",
  "public/js/vendor/bootstrap-responsive.css",
  "public/js/app.css",
))
~~~

Alternatively, plug your own `stylebundle.Options` into `stylebundle.Bundle(opts *stylebundle.Options, files ...string)` to customize your optimizations:

~~~ go
type Options struct {
  Minify bool
}
~~~