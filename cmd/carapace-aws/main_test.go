//go:build integration

package main

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-aws/cmd/carapace-aws/cmd/botocore"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
)

func TestService(t *testing.T) {
	serviceToTest := os.Getenv("SERVICE")
	for service := range botocore.Services() {
		t.Run(service, func(t *testing.T) {
			if serviceToTest != "" && service != serviceToTest { // TODO env var to only test a specific service
				t.SkipNow()
			}

			fmt.Printf("\033[2m# %v\033[0m\n", service)

			patch := carapace.DiffPatch(
				bridge.ActionAws("aws"),
				bridge.ActionCarapace("carapace-aws"),
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

			for operation := range botocore.Operations(service) {
				t.Run(operation, func(t *testing.T) {
					t.Parallel()
					patch := carapace.DiffPatch(
						bridge.ActionAws("aws"),
						bridge.ActionCarapace("carapace-aws"),
						carapace.NewContext(service, operation, "--"),
					)

					s := []string{fmt.Sprintf("\033[2m# %v %v\033[0m", service, operation)}
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
			}
		})
	}
}
