package main

// AppYaml as app.yaml
type AppYaml struct {
	Handlers []Route `yaml:"handlers"`
}
