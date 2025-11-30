package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	spec "github.com/carapace-sh/carapace-spec"
	"gopkg.in/yaml.v3"
)

func main() {
	content, err := os.ReadFile("../../../package.json")
	if err != nil {
		panic(err.Error())
	}

	var pkg struct {
		Dependencies map[string]string
	}
	if err := json.Unmarshal(content, &pkg); err != nil {
		panic(err.Error())
	}

	repo, tag, ok := strings.Cut(pkg.Dependencies["aws"], "#")
	if !ok {
		panic("failed to read aws dependency")
	}

	dir, err := os.MkdirTemp("", "carapace-aws_clone*")
	if err != nil {
		panic(err.Error())
	}
	defer os.RemoveAll(dir)

	if err := execCommand("git", "clone", repo, "--branch", tag, "--depth", "1", dir); err != nil {
		panic(err.Error())
	}

	// TODO remove existing specs first (e.g. removed/renamed service)
	if err := execCommand("go", "run", "../../carapace-spec-botocore",
		"--no-doc",
		"--target", "../cmd/botocore",
		filepath.Join(dir, "awscli/botocore/data")); err != nil {
		panic(err.Error())
	}

	services, err := services("../cmd/botocore")
	if err != nil {
		panic(err.Error())
	}
	err = os.WriteFile("../cmd/botocore/botocore_generated.go", []byte(services), os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
	err = execCommand("go", "fmt", "../cmd/botocore/botocore_generated.go")
	if err != nil {
		panic(err.Error())
	}
}

func execCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	println("# " + strings.Join(cmd.Args, " "))
	return cmd.Run()
}

func services(dir string) (string, error) {
	r := regexp.MustCompile(`^aws\.(?P<service>[^.]+)\.yaml$`)
	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}

	s := `package botocore

func init() {
	services = map[string]string{
`

	for _, entry := range entries {
		if !r.MatchString(entry.Name()) {
			continue
		}

		content, err := os.ReadFile(filepath.Join(dir, entry.Name()))
		if err != nil {
			return "", err
		}

		var command spec.Command
		if err := yaml.Unmarshal(content, &command); err != nil {
			return "", err
		}

		s += fmt.Sprintf("\t\t%q: %q,\n", command.Name, command.Description)
	}

	s += `	}
}
`
	return s, nil
}
