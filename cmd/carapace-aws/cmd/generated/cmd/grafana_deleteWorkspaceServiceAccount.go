package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var grafana_deleteWorkspaceServiceAccountCmd = &cobra.Command{
	Use:   "delete-workspace-service-account",
	Short: "Deletes a workspace service account from the workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(grafana_deleteWorkspaceServiceAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(grafana_deleteWorkspaceServiceAccountCmd).Standalone()

		grafana_deleteWorkspaceServiceAccountCmd.Flags().String("service-account-id", "", "The ID of the service account to delete.")
		grafana_deleteWorkspaceServiceAccountCmd.Flags().String("workspace-id", "", "The ID of the workspace where the service account resides.")
		grafana_deleteWorkspaceServiceAccountCmd.MarkFlagRequired("service-account-id")
		grafana_deleteWorkspaceServiceAccountCmd.MarkFlagRequired("workspace-id")
	})
	grafanaCmd.AddCommand(grafana_deleteWorkspaceServiceAccountCmd)
}
