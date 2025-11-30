package main

import "github.com/carapace-sh/carapace-aws/cmd/carapace-aws/cmd"

//go:generate sh -c "go run -C ./generate ."
func main() {
	cmd.Execute()
}
