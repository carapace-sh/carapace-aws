package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_updateConnectClientAddInCmd = &cobra.Command{
	Use:   "update-connect-client-add-in",
	Short: "Updates a Amazon Connect client add-in.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_updateConnectClientAddInCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_updateConnectClientAddInCmd).Standalone()

		workspaces_updateConnectClientAddInCmd.Flags().String("add-in-id", "", "The identifier of the client add-in to update.")
		workspaces_updateConnectClientAddInCmd.Flags().String("name", "", "The name of the client add-in.")
		workspaces_updateConnectClientAddInCmd.Flags().String("resource-id", "", "The directory identifier for which the client add-in is configured.")
		workspaces_updateConnectClientAddInCmd.Flags().String("url", "", "The endpoint URL of the Amazon Connect client add-in.")
		workspaces_updateConnectClientAddInCmd.MarkFlagRequired("add-in-id")
		workspaces_updateConnectClientAddInCmd.MarkFlagRequired("resource-id")
	})
	workspacesCmd.AddCommand(workspaces_updateConnectClientAddInCmd)
}
