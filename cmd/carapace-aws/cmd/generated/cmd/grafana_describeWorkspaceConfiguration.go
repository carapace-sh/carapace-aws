package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var grafana_describeWorkspaceConfigurationCmd = &cobra.Command{
	Use:   "describe-workspace-configuration",
	Short: "Gets the current configuration string for the given workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(grafana_describeWorkspaceConfigurationCmd).Standalone()

	grafana_describeWorkspaceConfigurationCmd.Flags().String("workspace-id", "", "The ID of the workspace to get configuration information for.")
	grafana_describeWorkspaceConfigurationCmd.MarkFlagRequired("workspace-id")
	grafanaCmd.AddCommand(grafana_describeWorkspaceConfigurationCmd)
}
