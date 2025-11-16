package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var grafana_createWorkspaceApiKeyCmd = &cobra.Command{
	Use:   "create-workspace-api-key",
	Short: "Creates a Grafana API key for the workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(grafana_createWorkspaceApiKeyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(grafana_createWorkspaceApiKeyCmd).Standalone()

		grafana_createWorkspaceApiKeyCmd.Flags().String("key-name", "", "Specifies the name of the key.")
		grafana_createWorkspaceApiKeyCmd.Flags().String("key-role", "", "Specifies the permission level of the key.")
		grafana_createWorkspaceApiKeyCmd.Flags().String("seconds-to-live", "", "Specifies the time in seconds until the key expires.")
		grafana_createWorkspaceApiKeyCmd.Flags().String("workspace-id", "", "The ID of the workspace to create an API key.")
		grafana_createWorkspaceApiKeyCmd.MarkFlagRequired("key-name")
		grafana_createWorkspaceApiKeyCmd.MarkFlagRequired("key-role")
		grafana_createWorkspaceApiKeyCmd.MarkFlagRequired("seconds-to-live")
		grafana_createWorkspaceApiKeyCmd.MarkFlagRequired("workspace-id")
	})
	grafanaCmd.AddCommand(grafana_createWorkspaceApiKeyCmd)
}
