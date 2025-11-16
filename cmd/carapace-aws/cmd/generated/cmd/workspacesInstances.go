package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesInstancesCmd = &cobra.Command{
	Use:   "workspaces-instances",
	Short: "Amazon WorkSpaces Instances provides an API framework for managing virtual workspace environments across multiple AWS regions, enabling programmatic creation and configuration of desktop infrastructure.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesInstancesCmd).Standalone()

	})
	rootCmd.AddCommand(workspacesInstancesCmd)
}
