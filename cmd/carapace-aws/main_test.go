//go:build integration

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-aws/cmd/carapace-aws/cmd/botocore"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace-spec/pkg/command"
)

func TestService(t *testing.T) {
	testDir, err := os.MkdirTemp("", "carapace-aws_testService-*")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer os.Remove(testDir)
	testFile, err := os.Create(filepath.Join(testDir, "outfile"))
	if err != nil {
		t.Fatal(err.Error())
	}
	defer os.Remove(testFile.Name())

	serviceToTest := os.Getenv("SERVICE")

	if serviceToTest == "" || serviceToTest == "_" {
		fmt.Printf("\033[2m# %v\033[0m\n", "_")

		patch := carapace.DiffPatch(
			bridge.ActionAws("aws"),
			bridge.ActionCarapace("carapace-aws").Chdir(testDir),
			carapace.NewContext(""),
		)
		for _, line := range patch {
			switch {
			case strings.HasPrefix(line, "-"):
				fmt.Printf("\033[0;31m%v\033[0m\n", line)
				t.Fail()
			case strings.HasPrefix(line, "+"):
				fmt.Printf("\033[0;32m%v\033[0m\n", line)
				t.Fail()
			}
		}
	}

	for service := range botocore.Services() {
		t.Run(service, func(t *testing.T) {
			if serviceToTest != "" && service != serviceToTest { // TODO env var to only test a specific service
				t.SkipNow()
			}

			fmt.Printf("\033[2m# %v\033[0m\n", service)

			patch := carapace.DiffPatch(
				bridge.ActionAws("aws"),
				bridge.ActionCarapace("carapace-aws").Chdir(testDir),
				carapace.NewContext(service, ""),
			)
			for _, line := range patch {
				switch {
				case strings.HasPrefix(line, "-"):
					fmt.Printf("\033[0;31m%v\033[0m\n", line)
					t.Fail()
				case strings.HasPrefix(line, "+"):
					fmt.Printf("\033[0;32m%v\033[0m\n", line)
					t.Fail()
				}
			}

			command, err := botocore.Get(fmt.Sprintf("aws.%s.yaml", service))
			if err != nil {
				t.Fatal(err.Error())
			}
			for _, operation := range command.Commands {
				testCommand(t, testDir, service, nil, operation)
			}
		})
	}
}

func testCommand(t *testing.T, testDir string, service string, parentPath []string, operation command.Command) {
	commandPath := append(append([]string{}, parentPath...), operation.Name)

	t.Run(strings.Join(commandPath, "/"), func(t *testing.T) {
		t.Parallel()

		baseContext := append([]string{service}, commandPath...)
		positionalContext := append(append([]string{}, baseContext...), "")
		flagContext := append(append([]string{}, baseContext...), "--")

		patch := carapace.DiffPatch(
			bridge.ActionAws("aws"),
			carapace.Batch(
				bridge.ActionCarapace("carapace-aws"),
				// force positional completion for outfile if available (returned by aws completer for streaming operations)
				bridge.ActionCarapace("carapace-aws").
					Chdir(testDir).
					Invoke(carapace.NewContext(positionalContext...)).
					Retain("outfile").ToA(),
			).ToA(),
			carapace.NewContext(flagContext...),
		)

		s := []string{fmt.Sprintf("\033[2m# %v %v\033[0m", service, strings.Join(commandPath, " "))}
		for _, line := range patch {
			switch {
			case strings.HasPrefix(line, "-"):
				s = append(s, fmt.Sprintf("\033[0;31m%v\033[0m", line))
				t.Fail()
			case strings.HasPrefix(line, "+"):
				s = append(s, fmt.Sprintf("\033[0;32m%v\033[0m", line))
				t.Fail()
			}
		}
		fmt.Println(strings.Join(s, "\n")) // TODO locking needed?
	})

	for _, subCommand := range operation.Commands {
		testCommand(t, testDir, service, commandPath, subCommand)
	}
}
