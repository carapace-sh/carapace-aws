package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_terminateWorkspacesCmd = &cobra.Command{
	Use:   "terminate-workspaces",
	Short: "Terminates the specified WorkSpaces.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_terminateWorkspacesCmd).Standalone()

	workspaces_terminateWorkspacesCmd.Flags().String("terminate-workspace-requests", "", "The WorkSpaces to terminate.")
	workspaces_terminateWorkspacesCmd.MarkFlagRequired("terminate-workspace-requests")
	workspacesCmd.AddCommand(workspaces_terminateWorkspacesCmd)
}
