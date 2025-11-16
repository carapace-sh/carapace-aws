package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_updateWorkspaceConfigurationCmd = &cobra.Command{
	Use:   "update-workspace-configuration",
	Short: "Use this operation to create or update the label sets, label set limits, and retention period of a workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_updateWorkspaceConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amp_updateWorkspaceConfigurationCmd).Standalone()

		amp_updateWorkspaceConfigurationCmd.Flags().String("client-token", "", "You can include a token in your operation to make it an idempotent opeartion.")
		amp_updateWorkspaceConfigurationCmd.Flags().String("limits-per-label-set", "", "This is an array of structures, where each structure defines a label set for the workspace, and defines the active time series limit for each of those label sets.")
		amp_updateWorkspaceConfigurationCmd.Flags().String("retention-period-in-days", "", "Specifies how many days that metrics will be retained in the workspace.")
		amp_updateWorkspaceConfigurationCmd.Flags().String("workspace-id", "", "The ID of the workspace that you want to update.")
		amp_updateWorkspaceConfigurationCmd.MarkFlagRequired("workspace-id")
	})
	ampCmd.AddCommand(amp_updateWorkspaceConfigurationCmd)
}
