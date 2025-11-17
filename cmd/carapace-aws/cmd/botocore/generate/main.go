package main

import (
	"fmt"
	"os"
	"regexp"

	spec "github.com/carapace-sh/carapace-spec"
	"gopkg.in/yaml.v3"
)

func main() {
	r := regexp.MustCompile(`^aws\.(?P<service>[^.]+)\.yaml$`)
	entries, err := os.ReadDir(".")
	if err != nil {
		panic(err.Error())
	}

	fmt.Print(`package botocore

func init() {
	services = map[string]string{
`)

	for _, entry := range entries {
		if !r.MatchString(entry.Name()) {
			continue
		}

		content, err := os.ReadFile(entry.Name())
		if err != nil {
			panic(err.Error())
		}

		var command spec.Command
		if err := yaml.Unmarshal(content, &command); err != nil {
			panic(err.Error())
		}

		fmt.Printf("\t\t%q: %q,\n", command.Name, command.Description)
	}

	fmt.Print(`	}
}
`)
}
