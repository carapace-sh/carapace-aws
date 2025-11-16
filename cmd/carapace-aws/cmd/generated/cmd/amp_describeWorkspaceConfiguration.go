package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_describeWorkspaceConfigurationCmd = &cobra.Command{
	Use:   "describe-workspace-configuration",
	Short: "Use this operation to return information about the configuration of a workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_describeWorkspaceConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amp_describeWorkspaceConfigurationCmd).Standalone()

		amp_describeWorkspaceConfigurationCmd.Flags().String("workspace-id", "", "The ID of the workspace that you want to retrieve information for.")
		amp_describeWorkspaceConfigurationCmd.MarkFlagRequired("workspace-id")
	})
	ampCmd.AddCommand(amp_describeWorkspaceConfigurationCmd)
}
