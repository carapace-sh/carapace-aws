package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var grafana_createWorkspaceServiceAccountCmd = &cobra.Command{
	Use:   "create-workspace-service-account",
	Short: "Creates a service account for the workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(grafana_createWorkspaceServiceAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(grafana_createWorkspaceServiceAccountCmd).Standalone()

		grafana_createWorkspaceServiceAccountCmd.Flags().String("grafana-role", "", "The permission level to use for this service account.")
		grafana_createWorkspaceServiceAccountCmd.Flags().String("name", "", "A name for the service account.")
		grafana_createWorkspaceServiceAccountCmd.Flags().String("workspace-id", "", "The ID of the workspace within which to create the service account.")
		grafana_createWorkspaceServiceAccountCmd.MarkFlagRequired("grafana-role")
		grafana_createWorkspaceServiceAccountCmd.MarkFlagRequired("name")
		grafana_createWorkspaceServiceAccountCmd.MarkFlagRequired("workspace-id")
	})
	grafanaCmd.AddCommand(grafana_createWorkspaceServiceAccountCmd)
}
