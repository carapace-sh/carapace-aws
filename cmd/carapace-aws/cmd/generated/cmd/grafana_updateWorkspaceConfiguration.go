package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var grafana_updateWorkspaceConfigurationCmd = &cobra.Command{
	Use:   "update-workspace-configuration",
	Short: "Updates the configuration string for the given workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(grafana_updateWorkspaceConfigurationCmd).Standalone()

	grafana_updateWorkspaceConfigurationCmd.Flags().String("configuration", "", "The new configuration string for the workspace.")
	grafana_updateWorkspaceConfigurationCmd.Flags().String("grafana-version", "", "Specifies the version of Grafana to support in the workspace.")
	grafana_updateWorkspaceConfigurationCmd.Flags().String("workspace-id", "", "The ID of the workspace to update.")
	grafana_updateWorkspaceConfigurationCmd.MarkFlagRequired("configuration")
	grafana_updateWorkspaceConfigurationCmd.MarkFlagRequired("workspace-id")
	grafanaCmd.AddCommand(grafana_updateWorkspaceConfigurationCmd)
}
