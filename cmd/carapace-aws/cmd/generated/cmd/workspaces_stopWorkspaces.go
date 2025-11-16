package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_stopWorkspacesCmd = &cobra.Command{
	Use:   "stop-workspaces",
	Short: "Stops the specified WorkSpaces.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_stopWorkspacesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_stopWorkspacesCmd).Standalone()

		workspaces_stopWorkspacesCmd.Flags().String("stop-workspace-requests", "", "The WorkSpaces to stop.")
		workspaces_stopWorkspacesCmd.MarkFlagRequired("stop-workspace-requests")
	})
	workspacesCmd.AddCommand(workspaces_stopWorkspacesCmd)
}
