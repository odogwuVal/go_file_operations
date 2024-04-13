package main

import (
	"fmt"

	"github.com/go-yaml/yaml"
)

var yamlFile = []byte(`
User: John Doak
ID: 25
Traits: ["Tall", "Blonde", "Dashing"]
`)

func main() {
	data := map[string]interface{}{}

	if err := yaml.Unmarshal(yamlFile, &data); err != nil {
		panic(err)
	}

	fmt.Println(data["User"])
	fmt.Println(data["ID"])

	// We check that this key is set so that our type assertion below will not panic.
	if _, ok := data["Traits"]; ok {
		// Because YAML can store lists of different types, all arrays are []interface{} when
		// decoding into a map.
		for _, trait := range data["Traits"].([]interface{}) {
			fmt.Println("Trait: ", trait)
		}
	}
}
