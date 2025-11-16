package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_getThingRuntimeConfigurationCmd = &cobra.Command{
	Use:   "get-thing-runtime-configuration",
	Short: "Get the runtime configuration of a thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_getThingRuntimeConfigurationCmd).Standalone()

	greengrass_getThingRuntimeConfigurationCmd.Flags().String("thing-name", "", "The thing name.")
	greengrass_getThingRuntimeConfigurationCmd.MarkFlagRequired("thing-name")
	greengrassCmd.AddCommand(greengrass_getThingRuntimeConfigurationCmd)
}
