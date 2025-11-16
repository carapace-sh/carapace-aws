package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_deleteConnectClientAddInCmd = &cobra.Command{
	Use:   "delete-connect-client-add-in",
	Short: "Deletes a client-add-in for Amazon Connect that is configured within a directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_deleteConnectClientAddInCmd).Standalone()

	workspaces_deleteConnectClientAddInCmd.Flags().String("add-in-id", "", "The identifier of the client add-in to delete.")
	workspaces_deleteConnectClientAddInCmd.Flags().String("resource-id", "", "The directory identifier for which the client add-in is configured.")
	workspaces_deleteConnectClientAddInCmd.MarkFlagRequired("add-in-id")
	workspaces_deleteConnectClientAddInCmd.MarkFlagRequired("resource-id")
	workspacesCmd.AddCommand(workspaces_deleteConnectClientAddInCmd)
}
