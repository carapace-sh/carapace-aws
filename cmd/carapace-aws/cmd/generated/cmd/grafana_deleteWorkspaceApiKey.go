package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var grafana_deleteWorkspaceApiKeyCmd = &cobra.Command{
	Use:   "delete-workspace-api-key",
	Short: "Deletes a Grafana API key for the workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(grafana_deleteWorkspaceApiKeyCmd).Standalone()

	grafana_deleteWorkspaceApiKeyCmd.Flags().String("key-name", "", "The name of the API key to delete.")
	grafana_deleteWorkspaceApiKeyCmd.Flags().String("workspace-id", "", "The ID of the workspace to delete.")
	grafana_deleteWorkspaceApiKeyCmd.MarkFlagRequired("key-name")
	grafana_deleteWorkspaceApiKeyCmd.MarkFlagRequired("workspace-id")
	grafanaCmd.AddCommand(grafana_deleteWorkspaceApiKeyCmd)
}
