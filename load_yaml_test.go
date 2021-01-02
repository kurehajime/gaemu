package main

import "testing"

func TestParseYaml(t *testing.T) {
	yamlString := `runtime: go114

automatic_scaling:
  min_idle_instances: 0
  max_instances: 2

handlers:
  - url: /colamone
    static_files: colamone
    upload: colamone
    expiration: "1d"
    secure: always

  - url: /(.*\.(gif|png|jpg|js|css|wav|zip|html|woff|woff2|map|ico))
    static_files: static/\1
    upload: static/(.*\.(gif|png|jpg|js|css|wav|zip|html|woff|woff2|map|ico))
    secure: always

  - url: /
    static_files: static/index.html
    upload: static/index.html
    secure: always
  
inbound_services:
  - warmup`
	result := parseYaml(yamlString)

	if len(result.Handlers) != 3 {
		t.Errorf("got: %v\nwant: %v", len(result.Handlers), 3)
	}

}
