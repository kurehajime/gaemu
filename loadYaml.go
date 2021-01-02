package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// LoadYaml : load app.yaml
func LoadYaml(path string) AppYaml {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return parseYaml(string(bytes))
}

func parseYaml(yamlString string) AppYaml {
	appYaml := AppYaml{}
	err := yaml.Unmarshal([]byte(yamlString), &appYaml)
	if err != nil {
		log.Fatal(err)
	}

	return appYaml
}
