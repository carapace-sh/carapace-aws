package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_startWorkspacesCmd = &cobra.Command{
	Use:   "start-workspaces",
	Short: "Starts the specified WorkSpaces.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_startWorkspacesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_startWorkspacesCmd).Standalone()

		workspaces_startWorkspacesCmd.Flags().String("start-workspace-requests", "", "The WorkSpaces to start.")
		workspaces_startWorkspacesCmd.MarkFlagRequired("start-workspace-requests")
	})
	workspacesCmd.AddCommand(workspaces_startWorkspacesCmd)
}
