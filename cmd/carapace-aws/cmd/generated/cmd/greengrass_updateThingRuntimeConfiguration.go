package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_updateThingRuntimeConfigurationCmd = &cobra.Command{
	Use:   "update-thing-runtime-configuration",
	Short: "Updates the runtime configuration of a thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_updateThingRuntimeConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_updateThingRuntimeConfigurationCmd).Standalone()

		greengrass_updateThingRuntimeConfigurationCmd.Flags().String("telemetry-configuration", "", "Configuration for telemetry service.")
		greengrass_updateThingRuntimeConfigurationCmd.Flags().String("thing-name", "", "The thing name.")
		greengrass_updateThingRuntimeConfigurationCmd.MarkFlagRequired("thing-name")
	})
	greengrassCmd.AddCommand(greengrass_updateThingRuntimeConfigurationCmd)
}
