package botocore

import (
	"embed"
	"regexp"

	spec "github.com/carapace-sh/carapace-spec"
	"gopkg.in/yaml.v3"
)

//go:embed *.yaml
var f embed.FS

//go:generate sh -c "go run ./generate > botocore_generated.go"
var services map[string]string

func Services() map[string]string {
	return services
}

func Operations(service string) map[string]string {
	entries, _ := f.ReadDir(".")
	r := regexp.MustCompile(`^aws\.(?P<service>[^.]+)\.(?P<operation>[^.]+)\.yaml$`)

	operations := make(map[string]string)
	for _, entry := range entries {
		if matches := r.FindStringSubmatch(entry.Name()); matches != nil {
			if matches[1] == service {
				operations[matches[2]] = "TODO"
			}
		}
	}
	return operations
}

func Get(name string) (*spec.Command, error) {
	content, err := f.ReadFile(name)
	if err != nil {
		return nil, err
	}

	var command *spec.Command
	if err := yaml.Unmarshal(content, &command); err != nil {
		return nil, err
	}
	return command, nil
}
