package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_terminateWorkspacesPoolSessionCmd = &cobra.Command{
	Use:   "terminate-workspaces-pool-session",
	Short: "Terminates the pool session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_terminateWorkspacesPoolSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_terminateWorkspacesPoolSessionCmd).Standalone()

		workspaces_terminateWorkspacesPoolSessionCmd.Flags().String("session-id", "", "The identifier of the pool session.")
		workspaces_terminateWorkspacesPoolSessionCmd.MarkFlagRequired("session-id")
	})
	workspacesCmd.AddCommand(workspaces_terminateWorkspacesPoolSessionCmd)
}
