package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var grafana_deleteWorkspaceServiceAccountTokenCmd = &cobra.Command{
	Use:   "delete-workspace-service-account-token",
	Short: "Deletes a token for the workspace service account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(grafana_deleteWorkspaceServiceAccountTokenCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(grafana_deleteWorkspaceServiceAccountTokenCmd).Standalone()

		grafana_deleteWorkspaceServiceAccountTokenCmd.Flags().String("service-account-id", "", "The ID of the service account from which to delete the token.")
		grafana_deleteWorkspaceServiceAccountTokenCmd.Flags().String("token-id", "", "The ID of the token to delete.")
		grafana_deleteWorkspaceServiceAccountTokenCmd.Flags().String("workspace-id", "", "The ID of the workspace from which to delete the token.")
		grafana_deleteWorkspaceServiceAccountTokenCmd.MarkFlagRequired("service-account-id")
		grafana_deleteWorkspaceServiceAccountTokenCmd.MarkFlagRequired("token-id")
		grafana_deleteWorkspaceServiceAccountTokenCmd.MarkFlagRequired("workspace-id")
	})
	grafanaCmd.AddCommand(grafana_deleteWorkspaceServiceAccountTokenCmd)
}
