package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_startWorkspacesPoolCmd = &cobra.Command{
	Use:   "start-workspaces-pool",
	Short: "Starts the specified pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_startWorkspacesPoolCmd).Standalone()

	workspaces_startWorkspacesPoolCmd.Flags().String("pool-id", "", "The identifier of the pool.")
	workspaces_startWorkspacesPoolCmd.MarkFlagRequired("pool-id")
	workspacesCmd.AddCommand(workspaces_startWorkspacesPoolCmd)
}
