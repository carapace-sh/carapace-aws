package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_restoreWorkspaceCmd = &cobra.Command{
	Use:   "restore-workspace",
	Short: "Restores the specified WorkSpace to its last known healthy state.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_restoreWorkspaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_restoreWorkspaceCmd).Standalone()

		workspaces_restoreWorkspaceCmd.Flags().String("workspace-id", "", "The identifier of the WorkSpace.")
		workspaces_restoreWorkspaceCmd.MarkFlagRequired("workspace-id")
	})
	workspacesCmd.AddCommand(workspaces_restoreWorkspaceCmd)
}
