package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_stopWorkspacesPoolCmd = &cobra.Command{
	Use:   "stop-workspaces-pool",
	Short: "Stops the specified pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_stopWorkspacesPoolCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_stopWorkspacesPoolCmd).Standalone()

		workspaces_stopWorkspacesPoolCmd.Flags().String("pool-id", "", "The identifier of the pool.")
		workspaces_stopWorkspacesPoolCmd.MarkFlagRequired("pool-id")
	})
	workspacesCmd.AddCommand(workspaces_stopWorkspacesPoolCmd)
}
