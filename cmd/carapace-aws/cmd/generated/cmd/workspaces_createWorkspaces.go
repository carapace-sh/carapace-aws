package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_createWorkspacesCmd = &cobra.Command{
	Use:   "create-workspaces",
	Short: "Creates one or more WorkSpaces.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_createWorkspacesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_createWorkspacesCmd).Standalone()

		workspaces_createWorkspacesCmd.Flags().String("workspaces", "", "The WorkSpaces to create.")
		workspaces_createWorkspacesCmd.MarkFlagRequired("workspaces")
	})
	workspacesCmd.AddCommand(workspaces_createWorkspacesCmd)
}
