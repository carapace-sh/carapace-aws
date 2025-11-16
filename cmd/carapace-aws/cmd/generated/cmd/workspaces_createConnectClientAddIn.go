package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_createConnectClientAddInCmd = &cobra.Command{
	Use:   "create-connect-client-add-in",
	Short: "Creates a client-add-in for Amazon Connect within a directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_createConnectClientAddInCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_createConnectClientAddInCmd).Standalone()

		workspaces_createConnectClientAddInCmd.Flags().String("name", "", "The name of the client add-in.")
		workspaces_createConnectClientAddInCmd.Flags().String("resource-id", "", "The directory identifier for which to configure the client add-in.")
		workspaces_createConnectClientAddInCmd.Flags().String("url", "", "The endpoint URL of the Amazon Connect client add-in.")
		workspaces_createConnectClientAddInCmd.MarkFlagRequired("name")
		workspaces_createConnectClientAddInCmd.MarkFlagRequired("resource-id")
		workspaces_createConnectClientAddInCmd.MarkFlagRequired("url")
	})
	workspacesCmd.AddCommand(workspaces_createConnectClientAddInCmd)
}
