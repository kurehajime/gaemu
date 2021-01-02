package main

// Route Static file or auto
type Route struct {
	URL         string `yaml:"url"`
	StaticFiles string `yaml:"static_files"`
	Staticdir   string `yaml:"static_dir"`
	Script      string `yaml:"script"`
}
