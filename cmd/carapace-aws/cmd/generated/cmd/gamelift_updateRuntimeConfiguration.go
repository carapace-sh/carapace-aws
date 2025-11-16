package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_updateRuntimeConfigurationCmd = &cobra.Command{
	Use:   "update-runtime-configuration",
	Short: "**This API works with the following fleet types:** EC2\n\nUpdates the runtime configuration for the specified fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_updateRuntimeConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_updateRuntimeConfigurationCmd).Standalone()

		gamelift_updateRuntimeConfigurationCmd.Flags().String("fleet-id", "", "A unique identifier for the fleet to update runtime configuration for.")
		gamelift_updateRuntimeConfigurationCmd.Flags().String("runtime-configuration", "", "Instructions for launching server processes on fleet computes.")
		gamelift_updateRuntimeConfigurationCmd.MarkFlagRequired("fleet-id")
		gamelift_updateRuntimeConfigurationCmd.MarkFlagRequired("runtime-configuration")
	})
	gameliftCmd.AddCommand(gamelift_updateRuntimeConfigurationCmd)
}
