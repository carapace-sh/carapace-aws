package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_updateWorkspacesPoolCmd = &cobra.Command{
	Use:   "update-workspaces-pool",
	Short: "Updates the specified pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_updateWorkspacesPoolCmd).Standalone()

	workspaces_updateWorkspacesPoolCmd.Flags().String("application-settings", "", "The persistent application settings for users in the pool.")
	workspaces_updateWorkspacesPoolCmd.Flags().String("bundle-id", "", "The identifier of the bundle.")
	workspaces_updateWorkspacesPoolCmd.Flags().String("capacity", "", "The desired capacity for the pool.")
	workspaces_updateWorkspacesPoolCmd.Flags().String("description", "", "Describes the specified pool to update.")
	workspaces_updateWorkspacesPoolCmd.Flags().String("directory-id", "", "The identifier of the directory.")
	workspaces_updateWorkspacesPoolCmd.Flags().String("pool-id", "", "The identifier of the specified pool to update.")
	workspaces_updateWorkspacesPoolCmd.Flags().String("running-mode", "", "The desired running mode for the pool.")
	workspaces_updateWorkspacesPoolCmd.Flags().String("timeout-settings", "", "Indicates the timeout settings of the specified pool.")
	workspaces_updateWorkspacesPoolCmd.MarkFlagRequired("pool-id")
	workspacesCmd.AddCommand(workspaces_updateWorkspacesPoolCmd)
}
