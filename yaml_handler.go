package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"path/filepath"
	"regexp"
)

// YamlHandler : Filehost or Reverse Proxy
type YamlHandler struct {
	path         string
	routes       []Route
	reverseProxy *httputil.ReverseProxy
}

// NewYamlHandler : return instance
func NewYamlHandler(reverseURL *url.URL, yamlPath string) *YamlHandler {
	y := new(YamlHandler)
	y.reverseProxy = httputil.NewSingleHostReverseProxy(reverseURL)
	y.path = filepath.Dir(yamlPath)
	appYaml := LoadYaml(yamlPath)
	y.routes = appYaml.Handlers
	return y
}

func (y *YamlHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	reDolToBackSlash := regexp.MustCompile(`\\(\d)`)
	for _, route := range y.routes {
		reg := regexp.MustCompile(route.URL)
		if reg.MatchString(r.URL.Path) {
			if route.Script == "auto" {
				y.reverseProxy.ServeHTTP(w, r)
			} else if route.Staticdir != "" {
				staticPath := reg.ReplaceAllString(r.URL.Path, reDolToBackSlash.ReplaceAllString(route.Staticdir, "$$$1"))
				http.ServeFile(w, r, path.Join(y.path, staticPath))
			} else {
				staticPath := reg.ReplaceAllString(r.URL.Path, reDolToBackSlash.ReplaceAllString(route.StaticFiles, "$$$1"))
				http.ServeFile(w, r, path.Join(y.path, staticPath))
			}
			return
		}
	}
	http.NotFound(w, r)
}
