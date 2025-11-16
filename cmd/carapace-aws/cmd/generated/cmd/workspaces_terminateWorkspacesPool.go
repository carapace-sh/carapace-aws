package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_terminateWorkspacesPoolCmd = &cobra.Command{
	Use:   "terminate-workspaces-pool",
	Short: "Terminates the specified pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_terminateWorkspacesPoolCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_terminateWorkspacesPoolCmd).Standalone()

		workspaces_terminateWorkspacesPoolCmd.Flags().String("pool-id", "", "The identifier of the pool.")
		workspaces_terminateWorkspacesPoolCmd.MarkFlagRequired("pool-id")
	})
	workspacesCmd.AddCommand(workspaces_terminateWorkspacesPoolCmd)
}
