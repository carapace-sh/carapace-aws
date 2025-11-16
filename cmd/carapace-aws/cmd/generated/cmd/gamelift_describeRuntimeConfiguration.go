package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_describeRuntimeConfigurationCmd = &cobra.Command{
	Use:   "describe-runtime-configuration",
	Short: "**This API works with the following fleet types:** EC2\n\nRetrieves a fleet's runtime configuration settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_describeRuntimeConfigurationCmd).Standalone()

	gamelift_describeRuntimeConfigurationCmd.Flags().String("fleet-id", "", "A unique identifier for the fleet to get the runtime configuration for.")
	gamelift_describeRuntimeConfigurationCmd.MarkFlagRequired("fleet-id")
	gameliftCmd.AddCommand(gamelift_describeRuntimeConfigurationCmd)
}
